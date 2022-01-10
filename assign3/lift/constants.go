package lift

const (
	Up                        int = 1
	Down                      int = 2
	Stop                      int = 0
	ErrNoMoreTargets              = Sentinel("no more targets to go")
	CurrentLevelNotAllowed        = Sentinel("not allowed to press the current level")
	TargetsFloorCannotBeEmpty     = Sentinel("target floor cannot be empty")
)
