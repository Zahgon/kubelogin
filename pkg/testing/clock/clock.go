package clock

import "time"

type Fake time.Time

func (f Fake) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
