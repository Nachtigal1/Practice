package integration_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/kre-college/lms/pkg/models"
	"github.com/kre-college/lms/pkg/user-management/repository/postgres"

	"github.com/golang-migrate/migrate/v4"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/stdlib"
	pg "github.com/kre-college/lms/pkg/db/postgres"
)

var testDBURL string
var db *pg.DB

var testTime, _ = time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
var testTimeData, _ = time.Parse("2006-01-02", time.Now().Format(time.RFC3339))

var countryCodes = []*models.CountryCodes{
	{"Afghanistan", "+93"}, {"Aland Islands", "+358"}, {"Albania", "+355"}, {"Algeria", "+213"}, {"AmericanSamoa", "+1684"}, {"Andorra", "+376"}, {"Angola", "+244"}, {"Anguilla", "+1264"}, {"Antarctica", "+672"}, {"Antigua and Barbuda", "+1268"}, {"Argentina", "+54"}, {"Armenia", "+374"}, {"Aruba", "+297"}, {"Australia", "+61"}, {"Austria", "+43"}, {"Azerbaijan", "+994"}, {"Bahamas", "+1242"}, {"Bahrain", "+973"}, {"Bangladesh", "+880"}, {"Barbados", "+1246"}, {"Belarus", "+375"}, {"Belgium", "+32"}, {"Belize", "+501"}, {"Benin", "+229"}, {"Bermuda", "+1441"}, {"Bhutan", "+975"}, {"Bolivia, Plurinational State of", "+591"}, {"Bosnia and Herzegovina", "+387"}, {"Botswana", "+267"}, {"Brazil", "+55"}, {"British Indian Ocean Territory", "+246"}, {"Brunei Darussalam", "+673"}, {"Bulgaria", "+359"}, {"Burkina Faso", "+226"}, {"Burundi", "+257"}, {"Cambodia", "+855"}, {"Cameroon", "+237"}, {"Canada", "+1"}, {"Cape Verde", "+238"}, {"Cayman Islands", "+ 345"}, {"Central African Republic", "+236"}, {"Chad", "+235"}, {"Chile", "+56"}, {"China", "+86"}, {"Christmas Island", "+61"}, {"Cocos (Keeling) Islands", "+61"}, {"Colombia", "+57"}, {"Comoros", "+269"}, {"Congo", "+242"}, {"Congo, The Democratic Republic of the Congo", "+243"}, {"Cook Islands", "+682"}, {"Costa Rica", "+506"}, {"Cote d`Ivoire", "+225"}, {"Croatia", "+385"}, {"Cuba", "+53"}, {"Cyprus", "+357"}, {"Czech Republic", "+420"}, {"Denmark", "+45"}, {"Djibouti", "+253"}, {"Dominica", "+1767"}, {"Dominican Republic", "+1849"}, {"Ecuador", "+593"}, {"Egypt", "+20"}, {"El Salvador", "+503"}, {"Equatorial Guinea", "+240"}, {"Eritrea", "+291"}, {"Estonia", "+372"}, {"Ethiopia", "+251"}, {"Falkland Islands (Malvinas)", "+500"}, {"Faroe Islands", "+298"}, {"Fiji", "+679"}, {"Finland", "+358"}, {"France", "+33"}, {"French Guiana", "+594"}, {"French Polynesia", "+689"}, {"Gabon", "+241"}, {"Gambia", "+220"}, {"Georgia", "+995"}, {"Germany", "+49"}, {"Ghana", "+233"}, {"Gibraltar", "+350"}, {"Greece", "+30"}, {"Greenland", "+299"}, {"Grenada", "+1473"}, {"Guadeloupe", "+590"}, {"Guam", "+1671"}, {"Guatemala", "+502"}, {"Guernsey", "+44"}, {"Guinea", "+224"}, {"Guinea-Bissau", "+245"}, {"Guyana", "+595"}, {"Haiti", "+509"}, {"Holy See (Vatican City State)", "+379"}, {"Honduras", "+504"}, {"Hong Kong", "+852"}, {"Hungary", "+36"}, {"Iceland", "+354"}, {"India", "+91"}, {"Indonesia", "+62"}, {"Iran, Islamic Republic of Persian Gulf", "+98"}, {"Iraq", "+964"}, {"Ireland", "+353"}, {"Isle of Man", "+44"}, {"Israel", "+972"}, {"Italy", "+39"}, {"Jamaica", "+1876"}, {"Japan", "+81"}, {"Jersey", "+44"}, {"Jordan", "+962"}, {"Kazakhstan", "+77"}, {"Kenya", "+254"}, {"Kiribati", "+686"}, {"Korea, Democratic People`s Republic of Korea", "+850"}, {"Korea, Republic of South Korea", "+82"}, {"Kuwait", "+965"}, {"Kyrgyzstan", "+996"}, {"Laos", "+856"}, {"Latvia", "+371"}, {"Lebanon", "+961"}, {"Lesotho", "+266"}, {"Liberia", "+231"}, {"Libyan Arab Jamahiriya", "+218"}, {"Liechtenstein", "+423"}, {"Lithuania", "+370"}, {"Luxembourg", "+352"}, {"Macao", "+853"}, {"Macedonia", "+389"}, {"Madagascar", "+261"}, {"Malawi", "+265"}, {"Malaysia", "+60"}, {"Maldives", "+960"}, {"Mali", "+223"}, {"Malta", "+356"}, {"Marshall Islands", "+692"}, {"Martinique", "+596"}, {"Mauritania", "+222"}, {"Mauritius", "+230"}, {"Mayotte", "+262"}, {"Mexico", "+52"}, {"Micronesia, Federated States of Micronesia", "+691"}, {"Moldova", "+373"}, {"Monaco", "+377"}, {"Mongolia", "+976"}, {"Montenegro", "+382"}, {"Montserrat", "+1664"}, {"Morocco", "+212"}, {"Mozambique", "+258"}, {"Myanmar", "+95"}, {"Namibia", "+264"}, {"Nauru", "+674"}, {"Nepal", "+977"}, {"Netherlands", "+31"}, {"Netherlands Antilles", "+599"}, {"New Caledonia", "+687"}, {"New Zealand", "+64"}, {"Nicaragua", "+505"}, {"Niger", "+227"}, {"Nigeria", "+234"}, {"Niue", "+683"}, {"Norfolk Island", "+672"}, {"Northern Mariana Islands", "+1670"}, {"Norway", "+47"}, {"Oman", "+968"}, {"Pakistan", "+92"}, {"Palau", "+680"}, {"Palestinian Territory, Occupied", "+970"}, {"Panama", "+507"}, {"Papua New Guinea", "+675"}, {"Paraguay", "+595"}, {"Peru", "+51"}, {"Philippines", "+63"}, {"Pitcairn", "+872"}, {"Poland", "+48"}, {"Portugal", "+351"}, {"Puerto Rico", "+1939"}, {"Qatar", "+974"}, {"Romania", "+40"}, {"Russia", "+7"}, {"Rwanda", "+250"}, {"Reunion", "+262"}, {"Saint Barthelemy", "+590"}, {"Saint Helena, Ascension and Tristan Da Cunha", "+290"}, {"Saint Kitts and Nevis", "+1869"}, {"Saint Lucia", "+1758"}, {"Saint Martin", "+590"}, {"Saint Pierre and Miquelon", "+508"}, {"Saint Vincent and the Grenadines", "+1784"}, {"Samoa", "+685"}, {"San Marino", "+378"}, {"Sao Tome and Principe", "+239"}, {"Saudi Arabia", "+966"}, {"Senegal", "+221"}, {"Serbia", "+381"}, {"Seychelles", "+248"}, {"Sierra Leone", "+232"}, {"Singapore", "+65"}, {"Slovakia", "+421"}, {"Slovenia", "+386"}, {"Solomon Islands", "+677"}, {"Somalia", "+252"}, {"South Africa", "+27"}, {"South Sudan", "+211"}, {"South Georgia and the South Sandwich Islands", "+500"}, {"Spain", "+34"}, {"Sri Lanka", "+94"}, {"Sudan", "+249"}, {"Suriname", "+597"}, {"Svalbard and Jan Mayen", "+47"}, {"Swaziland", "+268"}, {"Sweden", "+46"}, {"Switzerland", "+41"}, {"Syrian Arab Republic", "+963"}, {"Taiwan", "+886"}, {"Tajikistan", "+992"}, {"Tanzania, United Republic of Tanzania", "+255"}, {"Thailand", "+66"}, {"Timor-Leste", "+670"}, {"Togo", "+228"}, {"Tokelau", "+690"}, {"Tonga", "+676"}, {"Trinidad and Tobago", "+1868"}, {"Tunisia", "+216"}, {"Turkey", "+90"}, {"Turkmenistan", "+993"}, {"Turks and Caicos Islands", "+1649"}, {"Tuvalu", "+688"}, {"Uganda", "+256"}, {"Ukraine", "+380"}, {"United Arab Emirates", "+971"}, {"United Kingdom", "+44"}, {"United States", "+1"}, {"Uruguay", "+598"}, {"Uzbekistan", "+998"}, {"Vanuatu", "+678"}, {"Venezuela, Bolivarian Republic of Venezuela", "+58"}, {"Vietnam", "+84"}, {"Virgin Islands, British", "+1284"}, {"Virgin Islands, U.S.", "+1340"}, {"Wallis and Futuna", "+681"}, {"Yemen", "+967"}, {"Zambia", "+260"}, {"Zimbabwe", "+263"},
}

var salt = "7vjKaPQ1aH"
var hash = "L9gThZKe/g6Vmr/tEc4q85nv5FbaDBts0WuQrrkwwnQ="

var groupsIDs = []uint64{1}

var user = &models.User{
	UserID: 1,
	Photo: &models.Photo{
		UserID: 1,
		Photo:  nil,
	},
	ContactDetails: &models.ContactDetails{
		UserID:      1,
		Email:       "mainAdmin@gmail.com",
		PhoneNumber: "+380951573051",
		CreatedAt:   testTime,
		ModifiedAt:  nil,
	},
	DateOfBirth: "15.12.2002",
	Deactivated: false,
	FirstName:   "Bob",
	GroupID:     nil,
	LastName:    "Brown",
	ModifiedAt:  nil,
	Roles:       []string{"Administrator", "Curator"},
	CreatedAt:   testTime,
}

var secondUser = &models.User{
	UserID: 1,
	Photo: &models.Photo{
		UserID: 1,
		Photo:  nil,
	},
	ContactDetails: &models.ContactDetails{
		UserID:      1,
		Email:       "mainAdmin@gmail.com",
		PhoneNumber: "+380951573051",
		CreatedAt:   testTime,
		ModifiedAt:  nil,
	},
	DateOfBirth: "15.12.2002",
	Deactivated: false,
	FirstName:   "Bob",
	GroupID:     nil,
	LastName:    "Brown",
	ModifiedAt:  nil,
	Roles:       []string{"Administrator", "Curator"},
	CreatedAt:   testTime,
}

var emails = []string{user.ContactDetails.Email, secondUser.ContactDetails.Email}
var phoneNumbers = []string{user.ContactDetails.PhoneNumber, secondUser.ContactDetails.PhoneNumber}

var incorrectEmails = []string{"notFound@gmail.com", "notFound@i.ua"}

var users = []*models.User{
	user,
	secondUser,
}

var incorrectUser = &models.User{
	UserID: 0,
	Photo: &models.Photo{
		UserID: 0,
		Photo:  nil,
	},
	ContactDetails: &models.ContactDetails{
		UserID:      0,
		Email:       "",
		PhoneNumber: "",
		CreatedAt:   testTime,
		ModifiedAt:  &testTime,
	},
	DateOfBirth: "",
	Deactivated: false,
	FirstName:   "",
	GroupID:     nil,
	LastName:    "",
	ModifiedAt:  &testTime,
	Roles:       []string{"Administrator", "Curator"},
}

var credentials = &models.Credentials{
	UserID: 1000000,
	Salt:   &salt,
	Hash:   &hash,
}

var incorrectGroup = &models.Group{
	ID:                            0,
	Name:                          "56",
	CreatedAt:                     testTime,
	ModifiedAt:                    &testTime,
	AcademicYearID:                0,
	Course:                        2,
	StartOfFirstSemester:          testTimeData,
	EndOfFirstSemester:            testTimeData,
	NumberOfWeeksInFirstSemester:  18,
	StartOfSecondSemester:         testTimeData,
	EndOfSecondSemester:           testTimeData,
	NumberOfWeeksInSecondSemester: 17,
	NumberOfStudents:              24,
	CuratorID:                     29,
	SpecialtyID:                   4,
	DepartmentID:                  8,
	OnContractualBasis:            false,
	CycleCommitteeID:              5,
	GroupRoom:                     "102",
	DateOfCommencementOfStudies:   testTimeData,
	DateOfGraduation:              testTimeData,
	SignOfDivision:                true,
	NumberOfBudgetStudents:        22,
	NumberOfContractors:           9,
	NumberOfPreferentialSeats:     3,
	IsDeleted:                     false,
}

var group = &models.Group{
	ID:                            1,
	Name:                          "56",
	CreatedAt:                     testTime,
	ModifiedAt:                    nil,
	Course:                        2,
	StartOfFirstSemester:          testTimeData,
	EndOfFirstSemester:            testTimeData,
	NumberOfWeeksInFirstSemester:  18,
	StartOfSecondSemester:         testTimeData,
	EndOfSecondSemester:           testTimeData,
	NumberOfWeeksInSecondSemester: 17,
	NumberOfStudents:              24,
	CuratorID:                     29,
	SpecialtyID:                   4,
	DepartmentID:                  8,
	OnContractualBasis:            false,
	CycleCommitteeID:              5,
	GroupRoom:                     "102",
	DateOfCommencementOfStudies:   testTimeData,
	DateOfGraduation:              testTimeData,
	SignOfDivision:                true,
	NumberOfBudgetStudents:        22,
	NumberOfContractors:           9,
	NumberOfPreferentialSeats:     3,
	IsDeleted:                     false,
}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "",
		Env: []string{
			"POSTGRES_USER=dev",
			"POSTGRES_PASSWORD=12345",
			"POSTGRES_DB=user-management",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		testDBURL = fmt.Sprintf("postgres://dev:12345@%v/user-management?sslmode=disable", resource.GetHostPort("5432/tcp"))
		db, err = pg.NewDB(testDBURL)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	err = migrateScripts("../migrations", testDBURL)
	if err != nil {
		log.Print(err.Error())
		return
	}

	log.Printf("testDBURL: %s", testDBURL)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestAddUsers(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       []*models.User
		expectedErr error
	}{
		{
			name:  "OK",
			input: users,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.AddUsers(ctx, testCase.input)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestAddContactDetails(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       *models.ContactDetails
		expectedErr error
	}{
		{
			name:  "OK",
			input: user.ContactDetails,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.AddContactDetails(ctx, testCase.input)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestAddUsersHistories(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       []*models.User
		expectedErr error
	}{
		{
			name:  "OK",
			input: users,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.AddUsersHistories(ctx, testCase.input)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       uint64
		expect      *models.User
		expectedErr error
	}{
		{
			name:   "OK",
			input:  user.UserID,
			expect: user,
		},
		{
			name:        "Not found",
			input:       0,
			expectedErr: sql.ErrNoRows,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.GetUserByID(ctx, testCase.input)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestFetchUsers(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		expect      []*models.User
		expectedErr error
	}{
		{
			name: "OK",
			expect: []*models.User{
				user,
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.FetchUsers(ctx)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       string
		expect      *models.User
		expectedErr error
	}{
		{
			name:   "OK",
			input:  user.ContactDetails.Email,
			expect: user,
		},
		{
			name:        "Not found",
			input:       "notFound@gmail.com",
			expectedErr: sql.ErrNoRows,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.GetUserByEmail(ctx, testCase.input)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestFetchContactDetails(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name             string
		inputEmail       []string
		inputPhoneNumber []string
		expect           []*models.ContactDetails
		expectedErr      error
	}{
		{
			name:       "OK",
			inputEmail: emails,
			expect:     []*models.ContactDetails{},
		},
		{
			name:             "OK",
			inputPhoneNumber: phoneNumbers,
			expect:           []*models.ContactDetails{},
		},
		{
			name:       "Not found",
			inputEmail: incorrectEmails,
			expect:     []*models.ContactDetails{},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.FetchContactDetails(ctx, testCase.inputEmail, testCase.inputPhoneNumber)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestAddPhoto(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       *models.User
		expectedErr error
	}{
		{
			name:  "OK",
			input: user,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.AddPhoto(ctx, testCase.input)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestUpdateUserByID(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       *models.User
		expectedErr error
	}{
		{
			name:  "OK",
			input: user,
		},
		{
			name:        "Not found",
			input:       incorrectUser,
			expectedErr: sql.ErrNoRows,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.UpdateUserByID(ctx, testCase.input)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}

}

func TestUpdatePhotoByID(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       *models.User
		expectedErr error
	}{
		{
			name:  "OK",
			input: user,
		},
		{
			name:        "Not found",
			input:       incorrectUser,
			expectedErr: sql.ErrNoRows,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.UpdatePhotoByID(ctx, testCase.input)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestUpdateContactDetailsByID(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       *models.ContactDetails
		expectedErr error
	}{
		{
			name:  "OK",
			input: user.ContactDetails,
		},
		{
			name:        "Not found",
			input:       incorrectUser.ContactDetails,
			expectedErr: sql.ErrNoRows,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.UpdateContactDetailsByID(ctx, testCase.input)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestGetCredentialsByUserID(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       uint64
		expect      *models.Credentials
		expectedErr error
	}{
		{
			name:   "OK",
			input:  1000000,
			expect: credentials,
		},
		{
			name:        "Not found",
			input:       0,
			expectedErr: sql.ErrNoRows,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.GetCredentialsByUserID(ctx, testCase.input)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestAddGroup(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       []*models.Group
		expectedErr error
	}{
		{
			name:  "OK",
			input: []*models.Group{group},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.AddGroups(ctx, testCase.input)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestGetGroupByID(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       uint64
		expect      *models.Group
		expectedErr error
	}{
		{
			name:   "OK",
			input:  group.ID,
			expect: group,
		},
		{
			name:        "Not found",
			input:       0,
			expectedErr: sql.ErrNoRows,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.GetGroupByID(ctx, testCase.input)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestGetCuratorID(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       int
		expect      int
		expectedErr error
	}{
		{
			name:   "OK",
			input:  group.CuratorID,
			expect: group.CuratorID,
		},
		{
			name:        "Not found",
			input:       0,
			expectedErr: sql.ErrNoRows,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.GetCuratorID(ctx, testCase.input)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestGetGroupName(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       string
		expect      string
		expectedErr error
	}{
		{
			name:   "OK",
			input:  group.Name,
			expect: group.Name,
		},
		{
			name:        "Not found",
			input:       "",
			expectedErr: sql.ErrNoRows,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.GetGroupName(ctx, testCase.input)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestFetchGroups(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		expect      []*models.Group
		expectedErr error
	}{
		{
			name: "OK",
			expect: []*models.Group{
				group,
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.FetchGroups(ctx)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestUpdateGroupByID(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       *models.Group
		expectedErr error
	}{
		{
			name:  "OK",
			input: group,
		},
		{
			name:        "Not found",
			input:       incorrectGroup,
			expectedErr: sql.ErrNoRows,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.UpdateGroupByID(ctx, testCase.input)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestDeleteGroups(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		input       []uint64
		expectedErr error
	}{
		{
			name:  "OK",
			input: groupsIDs,
		},
		{
			name:        "Not found",
			input:       []uint64{},
			expectedErr: sql.ErrNoRows,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.DeleteGroups(ctx, testCase.input, testTime)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestFetchCountryCodes(t *testing.T) {
	repo := postgres.NewUserRepository(db)
	ctx := context.Background()

	testTable := []struct {
		name        string
		expect      []*models.CountryCodes
		expectedErr error
	}{
		{
			name:   "OK",
			expect: countryCodes,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := repo.FetchCountryCodes(ctx)

			assert.Equal(t, testCase.expect, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func migrateScripts(migrations, dbURL string) error {
	mgrt, err := migrate.New("file://"+migrations, dbURL)
	if err != nil {
		return errors.Errorf("Could not open migration sources: %s, err: %s", migrations, err.Error())
	}

	err = mgrt.Up()
	if err != nil {
		return errors.Errorf("Could not apply migrations: %s, err: %s", migrations, err.Error())
	}
	return nil
}
