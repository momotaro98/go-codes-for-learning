package main

func NewApplicationWithFOP(course Course, ops ...Option) *Application {
    a := Application{Course: course}
    for _, option := range ops {
        option(&a)
    }
    return &a
}

type Option func(*Application)

func WithSupport(flg bool) Option {
    return func(a *Application) {
        a.SubscribeSupportService = flg
    }
}

func WithMovie(flg bool) Option {
    return func(a *Application) {
        a.SubscribeMovieService = flg
    }
}

func WithBackupService(flg bool) Option {
    return func(a *Application) {
        a.SubscribeBackupService = flg
    }
}
