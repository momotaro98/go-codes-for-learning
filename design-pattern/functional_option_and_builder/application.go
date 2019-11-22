package main

type Course int

const (
    Basic Course = iota
    Premium
)

func (c Course) String() string {
    switch c {
    case Basic:
        return "Basic"
    case Premium:
        return "Premium"
    default:
        return ""
    }
}

type Application struct {
    Course Course
    SubscribeSupportService bool
    SubscribeMovieService bool
    SubscribeBackupService bool
}
