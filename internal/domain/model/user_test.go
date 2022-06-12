package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const layout = "2006-01-02"

func timeParse(t *testing.T, value string) time.Time {
	res, err := time.Parse(layout, value)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func TestAge(t *testing.T) {
	cases := []struct {
		birthdate string
		today     string
		expected  uint
	}{
		{
			birthdate: "2000-06-01",
			today:     "2000-05-01",
			expected:  0,
		},
		{
			birthdate: "2000-05-31",
			today:     "2000-06-01",
			expected:  0,
		},
		{
			birthdate: "2000-06-02",
			today:     "2001-06-01",
			expected:  0,
		},
		{
			birthdate: "2000-06-01",
			today:     "2001-06-01",
			expected:  1,
		},
	}
	for i, tc := range cases {
		birthdate := timeParse(t, tc.birthdate)
		today := timeParse(t, tc.today)
		age := Age(birthdate, today)
		assert.Equal(t, tc.expected, age, fmt.Sprintf("case %d failed", i))
	}
}

func TestUserIsSenior(t *testing.T) {
	cases := []struct {
		age      uint
		expected bool
	}{
		{
			age:      65,
			expected: true,
		},
		{
			age:      64,
			expected: false,
		},
	}
	for i, tc := range cases {
		actual := IsSenior(tc.age)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("case %d failed", i))
	}
}

func TestUserIsAdult(t *testing.T) {
	cases := []struct {
		age      uint
		expected bool
	}{
		{
			age:      65,
			expected: false,
		},
		{
			age:      64,
			expected: true,
		},
		{
			age:      12,
			expected: true,
		},
		{
			age:      11,
			expected: false,
		},
	}
	for i, tc := range cases {
		actual := IsAdult(tc.age)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("case %d failed", i))
	}
}

func TestUserIsChild(t *testing.T) {
	cases := []struct {
		age      uint
		expected bool
	}{
		{
			age:      12,
			expected: false,
		},
		{
			age:      11,
			expected: true,
		},
		{
			age:      3,
			expected: true,
		},
		{
			age:      2,
			expected: false,
		},
	}
	for i, tc := range cases {
		actual := IsChild(tc.age)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("case %d failed", i))
	}
}

func TestUserFullName(t *testing.T) {
	user := &User{
		FirstName: "world",
		LastName:  "hello",
	}
	actual := user.FullName()
	assert.Equal(t, "hello world", actual)
}

func TestUserListHasChild(t *testing.T) {
	cases := []struct {
		users    UserList
		today    time.Time
		expected bool
	}{
		{
			users:    UserList{},
			today:    timeParse(t, "2000-06-01"),
			expected: false,
		},
		{
			users: UserList{
				User{
					Birthdate: timeParse(t, "2000-06-01"),
				},
			},
			today:    timeParse(t, "2001-06-01"),
			expected: false,
		},
		{
			users: UserList{
				User{
					Birthdate: timeParse(t, "2000-06-01"),
				},
			},
			today:    timeParse(t, "2005-06-01"),
			expected: true,
		},
	}
	for i, tc := range cases {
		actual := tc.users.HasChild(tc.today)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("case %d failed", i))
	}
}
