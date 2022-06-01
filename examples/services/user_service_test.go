package services

import (
	"fmt"
	"test-flow-code-gen/examples/models"
	"test-flow-code-gen/testflow"
	"testing"
)

type UserServiceTest interface {
	testflow.TestProcess
	TestCreate(t *testing.T)
	TestUpdate(t *testing.T)
	TestDelete(t *testing.T)
	TestFind(t *testing.T)
}

type userServiceTest struct {
	userServiceTestTitle userServiceTestTitle
	userService          *UserService
	// add common parameters between test cases
}

type userServiceTestTitle struct {
	titleTestCreate string
	titleTestUpdate string
	titleTestDelete string
	titleTestFind   string
}

func NewUserServiceTestObject(
	testTitles userServiceTestTitle,
) UserServiceTest {
	return &userServiceTest{
		userService: NewUserService(),
		userServiceTestTitle: userServiceTestTitle{
			titleTestCreate: testTitles.titleTestCreate,
			titleTestUpdate: testTitles.titleTestUpdate,
			titleTestDelete: testTitles.titleTestDelete,
			titleTestFind:   testTitles.titleTestFind,
		},
		// add initialize common parameters between test cases
	}
}

func NewUserServiceTestTitle() userServiceTestTitle {
	return userServiceTestTitle{
		titleTestCreate: "TestCreate",
		titleTestUpdate: "TestUpdate",
		titleTestDelete: "TestDelete",
		titleTestFind:   "TestFind",
	}
}

func TestUserService(t *testing.T) {
	testTitles := NewUserServiceTestTitle()
	testInterface := NewUserServiceTestObject(testTitles)
	testDescribes := []testflow.Describe{
		{Title: testTitles.titleTestCreate, Test: testInterface.TestCreate},
		{Title: testTitles.titleTestUpdate, Test: testInterface.TestUpdate},
		{Title: testTitles.titleTestDelete, Test: testInterface.TestDelete},
		{Title: testTitles.titleTestFind, Test: testInterface.TestFind},
	}
	testflow.TestSuite(t, testInterface, testDescribes)
}

func (ut *userServiceTest) BeforeAll(t *testing.T) {}

func (ut *userServiceTest) BeforeTest(t *testing.T, testTitle string) {
	// setup process by test method case
	switch testTitle {
	case ut.userServiceTestTitle.titleTestCreate:
	case ut.userServiceTestTitle.titleTestUpdate:
	case ut.userServiceTestTitle.titleTestDelete:
	case ut.userServiceTestTitle.titleTestFind:
	}
}

func (ut *userServiceTest) BeforeCase(t *testing.T, testTitle string) {
	// setup process by test case
	switch testTitle {
	case ut.userServiceTestTitle.titleTestCreate:
	case ut.userServiceTestTitle.titleTestUpdate:
	case ut.userServiceTestTitle.titleTestDelete:
	case ut.userServiceTestTitle.titleTestFind:
	}
}

func (ut *userServiceTest) AfterAll(t *testing.T) {}

func (ut *userServiceTest) AfterTest(t *testing.T, testTitle string) {
	// teardown process by test method case
	switch testTitle {
	case ut.userServiceTestTitle.titleTestCreate:
	case ut.userServiceTestTitle.titleTestUpdate:
	case ut.userServiceTestTitle.titleTestDelete:
	case ut.userServiceTestTitle.titleTestFind:
	}
}

func (ut *userServiceTest) AfterCase(t *testing.T, testTitle string) {
	// teardown process by test case
	switch testTitle {
	case ut.userServiceTestTitle.titleTestCreate:
	case ut.userServiceTestTitle.titleTestUpdate:
	case ut.userServiceTestTitle.titleTestDelete:
	case ut.userServiceTestTitle.titleTestFind:
	}
}

func (ut *userServiceTest) TestCreate(t *testing.T) {

	type Argument struct {
		user models.User
	}

	type TestCase struct {
		testflow.It
		input       Argument
		expect      int
		expectError error
	}
	testCases := []TestCase{
		// TODO design test case
		// {
		// 	It: testflow.It{
		// 		SerialNum:   1,
		// 		Title:       "",
		// 		TestClass:   testflow.TestClassNormal, // testflow.TestClassSemiNormal | testflow.TestClassError
		// 		ExpectError: nil,
		// 	},
		// 	input: Argument{},
		// 	expect: nil,
		// },
		// add test case
	}
	for _, testCase := range testCases {
		t.Log(testCase.SerialNum, testCase.TestClass.String(), testCase.Title, testCase.input)
		ut.BeforeCase(t, ut.userServiceTestTitle.titleTestCreate)

		result, err := ut.userService.Create(testCase.input.user)
		fmt.Println(result)
		fmt.Println(err)
		//assert.Equal(t, testCase.expect, result)
		//assert.Equal(t, testCase.expectError, err)

		ut.AfterCase(t, ut.userServiceTestTitle.titleTestCreate)
	}
}

func (ut *userServiceTest) TestUpdate(t *testing.T) {

	type Argument struct {
		user models.User
	}

	type TestCase struct {
		testflow.It
		input       Argument
		expect      error
		expectError error
	}
	testCases := []TestCase{
		// TODO design test case
		// {
		// 	It: testflow.It{
		// 		SerialNum:   1,
		// 		Title:       "",
		// 		TestClass:   testflow.TestClassNormal, // testflow.TestClassSemiNormal | testflow.TestClassError
		// 		ExpectError: nil,
		// 	},
		// 	input: Argument{},
		// 	expect: nil,
		// },
		// add test case
	}
	for _, testCase := range testCases {
		t.Log(testCase.SerialNum, testCase.TestClass.String(), testCase.Title, testCase.input)
		ut.BeforeCase(t, ut.userServiceTestTitle.titleTestUpdate)

		err := ut.userService.Update(testCase.input.user)
		fmt.Println(err)
		//assert.Equal(t, testCase.expectError, err)

		ut.AfterCase(t, ut.userServiceTestTitle.titleTestUpdate)
	}
}

func (ut *userServiceTest) TestDelete(t *testing.T) {

	type TestCase struct {
		testflow.It

		expect      error
		expectError error
	}
	testCases := []TestCase{
		// TODO design test case
		// {
		// 	It: testflow.It{
		// 		SerialNum:   1,
		// 		Title:       "",
		// 		TestClass:   testflow.TestClassNormal, // testflow.TestClassSemiNormal | testflow.TestClassError
		// 		ExpectError: nil,
		// 	},
		// 	input: Argument{},
		// 	expect: nil,
		// },
		// add test case
	}
	for _, testCase := range testCases {
		t.Log(testCase.SerialNum, testCase.TestClass.String(), testCase.Title)
		ut.BeforeCase(t, ut.userServiceTestTitle.titleTestDelete)

		err := ut.userService.Delete()
		fmt.Println(err)
		//assert.Equal(t, testCase.expectError, err)

		ut.AfterCase(t, ut.userServiceTestTitle.titleTestDelete)
	}
}

func (ut *userServiceTest) TestFind(t *testing.T) {

	type TestCase struct {
		testflow.It

		expect      []models.User
		expectError error
	}
	testCases := []TestCase{
		// TODO design test case
		// {
		// 	It: testflow.It{
		// 		SerialNum:   1,
		// 		Title:       "",
		// 		TestClass:   testflow.TestClassNormal, // testflow.TestClassSemiNormal | testflow.TestClassError
		// 		ExpectError: nil,
		// 	},
		// 	input: Argument{},
		// 	expect: nil,
		// },
		// add test case
	}
	for _, testCase := range testCases {
		t.Log(testCase.SerialNum, testCase.TestClass.String(), testCase.Title)
		ut.BeforeCase(t, ut.userServiceTestTitle.titleTestFind)

		result, err := ut.userService.Find()
		fmt.Println(result)
		fmt.Println(err)
		//assert.Equal(t, testCase.expect, result)
		//assert.Equal(t, testCase.expectError, err)

		ut.AfterCase(t, ut.userServiceTestTitle.titleTestFind)
	}
}
