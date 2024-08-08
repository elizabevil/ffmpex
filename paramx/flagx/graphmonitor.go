package flagx

//graphmonitor

type ModeFlags = Flag
type MFlags = ModeFlags

const (
	ModeFlags_full       ModeFlags = "full"
	ModeFlags_compact    ModeFlags = "compact"
	ModeFlags_nozero     ModeFlags = "nozero"
	ModeFlags_noeof      ModeFlags = "noeof"
	ModeFlags_nodisabled ModeFlags = "nodisabled"
)

type FlagFlag Flag //Set = Flags which enable which stats are shown in video.
