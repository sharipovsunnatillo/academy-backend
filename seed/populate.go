package seed

import (
	"fmt"
	"github.com/jaswdr/faker/v2"
	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
	"github.com/sharipov/sunnatillo/academy-backend/pkg/enums"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func Populate(db *gorm.DB) {
	fake := faker.New()
	regions(db, fake)
	districts(db, fake)
	roles(db)
	permissions(db)
	subjects(db, fake)
	textBooks(db, fake)
	timeSlots(db)
	users(db, fake)
	trainingCenters(db, fake)
	branches(db, fake)
	rooms(db, fake)
	teacherInfos(db, fake)
	groups(db, fake)
	lessons(db, fake)
	tasks(db, fake)
	attendances(db)
	grades(db)

	fmt.Println("Population finished!!!")
}

func grades(db *gorm.DB) {
	if err := db.Exec("insert into grades(student_id, lesson_id, task_id, score)\nselect gs.user_id, l.id, t.id, t.max_score * random()\nfrom tasks t\n         left join lessons l on t.lesson_id = l.id\n         left join group_students gs on gs.group_id = l.group_id").Error; err != nil {
		panic(err)
	}
}

func attendances(db *gorm.DB) {
	if err := db.Exec("insert into attendances(lesson_id, student_id, status)\nselect l.id, gs.user_id, case when random() > 0.3 then 'present' when random() > 0.7 then 'late' else 'absent' end\nfrom lessons l\n         left join group_students gs on gs.user_id = l.group_id").Error; err != nil {
		panic(err)
	}

}

func tasks(db *gorm.DB, fake faker.Faker) {
	var lessonIds []uint
	if err := db.Raw("select id from lessons order by random() limit (select count(id) /15 from lessons)").Scan(&lessonIds).Error; err != nil {
		panic(err)
	}
	var tasks []models.Task
	for _, id := range lessonIds {
		for i := 0; i < fake.IntBetween(1, 3); i++ {
			tasks = append(tasks, models.Task{
				Name:     fake.Lorem().Word(),
				LessonID: id,
				MaxScore: fake.Float32(0, 5, 100),
			})
		}
	}
	if err := db.CreateInBatches(tasks, 100).Error; err != nil {
		panic(err)
	}
	var documents []models.Document
	for _, task := range tasks {
		for i := 0; i < fake.IntBetween(1, 3); i++ {
			documents = append(documents, models.Document{
				Name:        fake.Lorem().Word(),
				Url:         fake.Internet().URL(),
				Type:        enums.Task,
				ReferenceID: task.ID,
				Reference:   reflect.TypeOf(models.Task{}).Name(),
			})
		}
	}
	if err := db.CreateInBatches(documents, 100).Error; err != nil {
		panic(err)
	}
}

func lessons(db *gorm.DB, fake faker.Faker) {
	type groupTimeSlotRow struct {
		ID         uint
		SubjectID  uint
		TimeSlotID uint
		BranchID   uint
		TeacherID  uint
		RoomID     uint
	}
	var groupTimeSlots []groupTimeSlotRow
	if err := db.Raw(`select g.id, g.subject_id, gt.time_slot_id, g.branch_id, g.teacher_id, g.room_id
from groups g
         left join group_timeslots gt on g.id = gt.group_id`).Find(&groupTimeSlots).Error; err != nil {
		panic(err)
	}
	lessonTypes := []enums.LessonType{enums.Lecture, enums.Lab, enums.Exam, enums.Exercise, enums.Test, enums.Other}
	var lessons []models.Lesson
	for i := 0; i < 90; i++ {
		for _, timeSlot := range groupTimeSlots {
			if len(lessons) == 1000 {
				if err := db.Create(lessons).Error; err != nil {
					panic(err)
				}
				lessons = lessons[:0]
			}
			lessons = append(lessons, models.Lesson{
				Name:       fake.Lorem().Sentence(fake.IntBetween(1, 3)),
				GroupID:    timeSlot.ID,
				SubjectID:  timeSlot.SubjectID,
				TimeSlotID: timeSlot.TimeSlotID,
				BranchID:   timeSlot.BranchID,
				TeacherID:  timeSlot.TeacherID,
				RoomID:     timeSlot.RoomID,
				Date:       time.Now().Add(time.Duration(i-45) * time.Hour * 24),
				Type:       lessonTypes[fake.IntBetween(0, len(lessonTypes)-1)],
			})
		}
	}
	if len(lessons) > 0 {
		if err := db.Create(lessons).Error; err != nil {
			panic(err)
		}
	}
}

func groups(db *gorm.DB, fake faker.Faker) {
	var subjects []models.Subject
	if err := db.Select("ID").Find(&subjects).Error; err != nil {
		panic(err)
	}
	var timeSlots []models.TimeSlot
	if err := db.Select("ID").Find(&timeSlots).Error; err != nil {
		panic(err)
	}
	var users []models.User
	if err := db.Select("ID").Find(&users).Error; err != nil {
		panic(err)
	}
	durations := []enums.TimeDuration{enums.Week, enums.Month}
	type branchTeacherSubject struct {
		BranchID  uint
		TeacherID uint
		RoomID    uint
	}
	var branches []branchTeacherSubject
	if err := db.Raw(`select b.id as branch_id, tb.teacher_info_id as teacher_id, r.id as room_id
from branches b
         inner join teacher_branches tb on tb.branch_id = b.id
         inner join rooms r on r.branch_id = b.id`).Scan(&branches).Error; err != nil {
		panic(err)
	}
	var groups []models.Group
	now := time.Now()
	for _, branch := range branches {
		for i := 0; i < fake.IntBetween(1, 4); i++ {
			timeSlot := timeSlots[fake.IntBetween(0, len(timeSlots)-1)]
			groups = append(groups, models.Group{
				Name:      fake.Lorem().Word(),
				BranchID:  branch.BranchID,
				TeacherID: branch.TeacherID,
				SubjectID: subjects[fake.IntBetween(0, len(subjects)-1)].ID,
				RoomID:    branch.RoomID,
				Started:   time.Date(now.Year(), now.Month(), now.Day(), timeSlot.Start.Hour(), timeSlot.Start.Minute(), now.Second(), 0, time.UTC),
				Ended:     time.Date(now.Year(), now.Month(), now.Day(), timeSlot.Start.Hour(), timeSlot.Start.Minute(), now.Second(), 0, time.UTC),
				Duration:  durations[fake.IntBetween(0, len(durations)-1)],
				Period:    fake.UInt16Between(4, 24),
			})
		}
	}
	if err := db.CreateInBatches(groups, 100).Error; err != nil {
		panic(err)
	}
	type groupStudentRow struct {
		GroupID uint
		UserID  uint
	}
	type groupTimeSlotRow struct {
		GroupID    uint
		TimeSlotId uint
	}
	var groupStudents []groupStudentRow
	var groupTimeSlots []groupTimeSlotRow
	for _, group := range groups {
		for i := 0; i < fake.IntBetween(2, 4); i++ {
			if len(groupTimeSlots) == 1000 {
				if err := db.Table("group_timeslots").Clauses(clause.OnConflict{DoNothing: true}).Create(groupTimeSlots).Error; err != nil {
					panic(err)
				}
				groupTimeSlots = groupTimeSlots[:0]
			}
			groupTimeSlots = append(groupTimeSlots, groupTimeSlotRow{
				GroupID:    group.ID,
				TimeSlotId: timeSlots[fake.IntBetween(0, len(timeSlots)-1)].ID,
			})
		}
		for i := 0; i < fake.IntBetween(4, 20); i++ {
			if len(groupStudents) == 1000 {
				if err := db.Table("group_students").Clauses(clause.OnConflict{DoNothing: true}).Create(groupStudents).Error; err != nil {
					panic(err)
				}
				groupStudents = groupStudents[:0]
			}
			groupStudents = append(groupStudents, groupStudentRow{
				GroupID: group.ID,
				UserID:  users[fake.IntBetween(0, len(users)-1)].ID,
			})
		}
	}
	if len(groupStudents) > 0 {
		if err := db.Table("group_students").Clauses(clause.OnConflict{DoNothing: true}).Create(groupStudents).Error; err != nil {
			panic(err)
		}
	}
	if len(groupTimeSlots) > 0 {
		if err := db.Table("group_timeslots").Clauses(clause.OnConflict{DoNothing: true}).Create(groupTimeSlots).Error; err != nil {
			panic(err)
		}
	}
}

func teacherInfos(db *gorm.DB, fake faker.Faker) {
	var users []models.User
	if err := db.Select("ID").Find(&users).Error; err != nil {
		panic(err)
	}
	teacherInfos := make([]models.TeacherInfo, fake.IntBetween(len(users)/5, len(users)))
	for i := 0; i < cap(teacherInfos); i++ {
		teacherInfos[i] = models.TeacherInfo{
			UserID: users[fake.IntBetween(0, len(users)-1)].ID,
		}
	}
	if err := db.CreateInBatches(teacherInfos, 100).Error; err != nil {
		panic(err)
	}
	//if err := db.Select("ID").Find(&teacherInfos).Error; err != nil {
	//	panic(err)
	//}
	type teacherSubjectsRow struct {
		TeacherInfoID uint
		SubjectID     uint
	}
	type teacherBranchesRow struct {
		TeacherInfoID uint
		BranchID      uint
	}
	var subjects []models.Subject
	var branches []models.Branch
	if err := db.Select("ID").Find(&subjects).Error; err != nil {
		panic(err)
	}
	if err := db.Select("ID").Find(&branches).Error; err != nil {
		panic(err)
	}
	documentTypes := []enums.DocumentType{enums.Passport, enums.IDCard, enums.DriverLicense, enums.Certificate, enums.ELSE}
	var teacherSubjects []teacherSubjectsRow
	var teacherBranches []teacherBranchesRow
	var documents []models.Document
	for _, teacherInfo := range teacherInfos {
		for i := 0; i < fake.IntBetween(1, 3); i++ {
			teacherSubjects = append(teacherSubjects, teacherSubjectsRow{
				TeacherInfoID: teacherInfo.ID,
				SubjectID:     subjects[fake.IntBetween(0, len(subjects)-1)].ID,
			})
		}
		for i := 0; i < fake.IntBetween(1, 3); i++ {
			teacherBranches = append(teacherBranches, teacherBranchesRow{
				TeacherInfoID: teacherInfo.ID,
				BranchID:      branches[fake.IntBetween(0, len(branches)-1)].ID,
			})
		}
		for _, documentType := range shuffleAndCut(documentTypes) {
			documents = append(documents, models.Document{
				Name:        fake.Lorem().Word(),
				Url:         fake.Internet().URL(),
				Type:        documentType,
				ReferenceID: teacherInfo.ID,
				Reference:   reflect.TypeOf(teacherInfo).Name(),
			})
		}
	}
	if err := db.Table("teacher_subjects").Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(teacherSubjects, 500).Error; err != nil {
		panic(err)
	}
	if err := db.Table("teacher_branches").Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(teacherBranches, 500).Error; err != nil {
		panic(err)
	}
	if err := db.CreateInBatches(documents, 100).Error; err != nil {
		panic(err)
	}
}

func rooms(db *gorm.DB, fake faker.Faker) {
	var branches []models.Branch
	if err := db.Select("ID").Find(&branches).Error; err != nil {
		panic(err)
	}
	rooms := make([]models.Room, fake.IntBetween(len(branches), len(branches)*3))
	for i := 0; i < cap(rooms); i++ {
		rooms[i] = models.Room{
			Number:   strconv.Itoa(fake.IntBetween(1, 100)),
			BranchID: branches[fake.IntBetween(0, len(branches)-1)].ID,
		}
	}
	if err := db.CreateInBatches(rooms, 100).Error; err != nil {
		panic(err)
	}

}

func branches(db *gorm.DB, fake faker.Faker) {
	var regions []models.Region
	var districts []models.District
	if err := db.Select("ID").Find(&regions).Error; err != nil {
		panic(err)
	}
	if err := db.Select("ID").Find(&districts).Error; err != nil {
		panic(err)
	}
	var centers []models.TrainingCenter
	if err := db.Select("ID").Find(&centers).Error; err != nil {
		panic(err)
	}
	branches := make([]models.Branch, fake.IntBetween(len(centers), len(centers)*2))
	for i := 0; i < cap(branches); i++ {
		branches[i] = models.Branch{
			Name:     fake.Address().City(),
			CenterID: centers[fake.IntBetween(0, len(centers)-1)].ID,
			Address: models.Address{
				RegionID:   regions[fake.IntBetween(0, len(regions)-1)].ID,
				DistrictID: districts[fake.IntBetween(0, len(districts)-1)].ID,
				Street:     fake.Address().StreetAddress(),
				House:      fake.Address().BuildingNumber(),
				Apartment:  strconv.Itoa(fake.IntBetween(1, 100)),
				Guide:      fake.Address().StreetAddress(),
				Longitude:  fake.Address().Longitude(),
				Latitude:   fake.Address().Latitude(),
			},
		}
	}
	if err := db.CreateInBatches(branches, 100).Error; err != nil {
		panic(err)
	}

}

func trainingCenters(db *gorm.DB, fake faker.Faker) {
	var regions []models.Region
	var districts []models.District
	if err := db.Select("ID").Find(&regions).Error; err != nil {
		panic(err)
	}
	if err := db.Select("ID").Find(&districts).Error; err != nil {
		panic(err)
	}
	trainingCenters := make([]models.TrainingCenter, fake.IntBetween(10, 50))
	for i := 0; i < cap(trainingCenters); i++ {
		trainingCenters[i] = models.TrainingCenter{
			Name: fake.Address().City(),
			Address: models.Address{
				RegionID:   regions[fake.IntBetween(0, len(regions)-1)].ID,
				DistrictID: districts[fake.IntBetween(0, len(districts)-1)].ID,
				Street:     fake.Address().StreetAddress(),
				House:      fake.Address().BuildingNumber(),
				Apartment:  strconv.Itoa(fake.IntBetween(1, 100)),
				Guide:      fake.Address().StreetAddress(),
				Longitude:  fake.Address().Longitude(),
				Latitude:   fake.Address().Latitude(),
			},
		}
	}
	if err := db.CreateInBatches(trainingCenters, 100).Error; err != nil {
		panic(err)
	}
}

func users(db *gorm.DB, fake faker.Faker) {
	var roles []models.Role
	var permissions []models.Permission
	err := db.Find(&roles).Error
	if err != nil {
		panic(err)
	}
	err = db.Find(&permissions).Error
	if err != nil {
		panic(err)
	}
	users := make([]models.User, fake.IntBetween(100, 500))
	for i := 0; i < cap(users); i++ {
		users[i] = models.User{
			FirstName:  fake.Person().FirstName(),
			LastName:   fake.Person().LastName(),
			MiddleName: fake.Person().FirstName(),
			Email:      fake.Internet().Email(),
			Phone:      fake.Phone().Number(),
			Password:   fake.RandomStringWithLength(fake.IntBetween(5, 255)),
		}
	}
	err = db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(users, 100).Error
	if err != nil {
		panic(err)
	}
	db.Select("ID").Find(&users)

	type userRolesRow struct {
		UserID   uint
		RoleName enums.Role
	}
	type userPermissionsRow struct {
		UserID         uint
		PermissionName string
	}

	var userRoles []userRolesRow
	var userPermissions []userPermissionsRow

	for _, user := range users {
		shuffledRoles := shuffleAndCut(roles)
		shuffledPermissions := shuffleAndCut(permissions)
		for _, role := range shuffledRoles {
			userRoles = append(userRoles, userRolesRow{
				UserID:   user.ID,
				RoleName: role.Name,
			})
		}
		for _, permission := range shuffledPermissions {
			userPermissions = append(userPermissions, userPermissionsRow{
				UserID:         user.ID,
				PermissionName: permission.Name,
			})
		}
	}
	if err := db.Table("user_roles").CreateInBatches(userRoles, 200).Error; err != nil {
		panic(err)
	}
	if err := db.Table("user_permissions").CreateInBatches(userPermissions, 200).Error; err != nil {
		panic(err)
	}

}

func timeSlots(db *gorm.DB) {
	timeSlots := make([]models.TimeSlot, 7*24)
	for i := 0; i < 7; i++ {
		for j := 0; j < 24; j++ {
			timeSlots[i*24+j] = models.TimeSlot{
				DayOfWeek: uint8(i),
				Start:     time.Date(1, 1, 1, j, 0, 0, 0, time.UTC),
				End:       time.Date(1, 1, 1, j+1, 0, 0, 0, time.UTC),
			}
		}
	}
	if err := db.CreateInBatches(timeSlots, 100).Error; err != nil {
		panic(err)
	}
}

func textBooks(db *gorm.DB, fake faker.Faker) {
	var subjects []models.Subject
	if err := db.Select("ID").Find(&subjects).Error; err != nil {
		panic(err)
	}
	textBooks := make([]models.TextBook, fake.IntBetween(len(subjects), len(subjects)*5))
	for i := 0; i < cap(textBooks); i++ {
		textBooks[i] = models.TextBook{
			Name:      fake.Lorem().Word(),
			Authors:   fake.Lorem().Sentence(fake.IntBetween(1, 5)),
			SubjectID: subjects[fake.IntBetween(0, len(subjects)-1)].ID,
		}
	}
	if err := db.CreateInBatches(textBooks, 100).Error; err != nil {
		panic(err)
	}
}

func subjects(db *gorm.DB, fake faker.Faker) {
	subjects := make([]models.Subject, fake.IntBetween(5, 10))
	for i := 0; i < cap(subjects); i++ {
		subjects[i] = models.Subject{
			Name: fake.Lorem().Word(),
		}
	}
	if err := db.Create(subjects).Error; err != nil {
		panic(err)
	}
}

func permissions(db *gorm.DB) {
	permissions := []models.Permission{
		{"read"},
		{"write"},
		{"delete"},
		{"update"},
		{"create"},
		{"execute"},
		{"manage"},
	}
	if err := db.Create(permissions).Error; err != nil {
		panic(err)
	}

}

func roles(db *gorm.DB) {
	roles := []models.Role{
		{enums.SUPER_ADMIN},
		{enums.ADMIN},
		{enums.TEACHER},
		{enums.STUDENT},
		{enums.GUEST},
		{enums.PARENT},
	}
	if err := db.Create(roles).Error; err != nil {
		panic(err)
	}

}

func districts(db *gorm.DB, fake faker.Faker) {
	var regions []models.Region
	err := db.Select("ID").Find(&regions).Error
	if err != nil {
		panic(err)
	}
	for _, region := range regions {
		i2 := fake.IntBetween(10, 20)
		districts := make([]models.District, i2)
		for i := 0; i < i2; i++ {
			districts[i] = models.District{
				Name:     fake.Address().City(),
				RegionID: region.ID,
			}
		}
		if err := db.CreateInBatches(districts, 100).Error; err != nil {
			panic(err)
		}
	}
}

func regions(db *gorm.DB, fake faker.Faker) {
	regionCount := fake.IntBetween(10, 20)
	regions := make([]models.Region, regionCount)
	for i := 0; i < regionCount; i++ {
		regions[i] = models.Region{
			Name: fake.Address().Country(),
		}
	}
	if err := db.Create(regions).Error; err != nil {
		panic(err)
	}
}

func shuffleAndCut[T any](input []T) []T {
	r := rand.New(rand.NewSource(time.Now().Unix())) // you could also pass in a *faker.Faker if you want reproducible seeding
	r.Shuffle(len(input), func(i, j int) {
		input[i], input[j] = input[j], input[i]
	})
	n := len(input)
	cutLen := r.Intn(n) + 1
	return input[:cutLen]
}
