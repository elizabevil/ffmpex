package muxerx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// 21.68 segment, stream_segment, ssegment
type Segmenter struct {
	IncrementTc typex.Bool `json:"increment_tc" flag:"-increment_tc"`
	//if set to 1, increment timecode between each segment If this is selected, the input need to have a timecode in the first video stream. Default value is 0.

	ReferenceStream typex.Specifier `json:"reference_stream" flag:"-reference_stream"`
	//Set the reference stream, as specified by the string specifier. If specifier is set to auto, the reference is chosen automatically. Otherwise it must be a stream specifier (see the “Stream specifiers” chapter in the ffmpeg manual) which specifies the reference stream. The default value is auto.

	SegmentFormat string `json:"segment_format" flag:"-segment_format"`
	//Override the inner container format, by default it is guessed by the filename extension.

	SegmentFormatOptions typex.OptionsList `json:"segment_format_options" flag:"-segment_format_options"`
	//Set output format options using a :-separated list of key=value parameters. Values containing the : special character must be escaped.

	SegmentList string `json:"segment_list" flag:"-segment_list"`
	//Generate also a listfile named name. If not specified no listfile is generated.

	SegmentListFlags flagx.SegmentFormatOptions `json:"segment_list_flags" flag:"-segment_list_flags"`

	SegmentListSize typex.Size `json:"segment_list_size" flag:"-segment_list_size"`
	//Update the list file so that it contains at most size segments. If 0 the list file will contain all the segments. Default value is 0.

	SegmentListEntryPrefix typex.Prefix `json:"segment_list_entry_prefix" flag:"-segment_list_entry_prefix"`
	//Prepend prefix to each entry. Useful to generate absolute paths. By default no prefix is applied.

	SegmentListType flagx.SegmentListType `json:"segment_list_type" flag:"-segment_list_type"`
	//Select the listing format.

	SegmentTime typex.Time `json:"segment_time" flag:"-segment_time"`
	//Set segment duration to time, the value must be a duration specification. Default value is "2". See also the segment_times option.

	//Note that splitting may not be accurate, unless you force the reference stream key-frames at the given time. See the introductory notice and the examples below.

	MinSegDuration typex.Time `json:"min_seg_duration" flag:"-min_seg_duration"`
	//Set minimum segment duration to time, the value must be a duration specification. This prevents the muxer ending segments at a duration below this value. Only effective with segment_time. Default value is "0".

	SegmentAtclocktime typex.Bool `json:"segment_atclocktime" flag:"-segment_atclocktime"`
	//If set to "1" split at regular clock time intervals starting from 00:00 o’clock. The time value specified in segment_time is used for setting the length of the splitting interval.
	//For example with segment_time set to "900" this makes it possible to create files at 12:00 o’clock, 12:15, 12:30, etc.
	//Default value is "0".

	SegmentClocktimeOffset typex.DurationI `json:"segment_clocktime_offset" flag:"-segment_clocktime_offset"`
	//Delay the segment splitting times with the specified duration when using segment_atclocktime.
	//For example with segment_time set to "900" and segment_clocktime_offset set to "300" this makes it possible to create files at 12:05, 12:20, 12:35, etc.
	//Default value is "0".

	SegmentClocktimeWrapDuration typex.DurationI `json:"segment_clocktime_wrap_duration" flag:"-segment_clocktime_wrap_duration"`
	//Force the segmenter to only start a new segment if a packet reaches the muxer within the specified duration after the segmenting clock time. This way you can make the segmenter more resilient to backward local time jumps, such as leap seconds or transition to standard time from daylight savings time.
	//Default is the maximum possible duration which means starting a new segment regardless of the elapsed time since the last clock time.

	SegmentTimeDelta typex.Delta `json:"segment_time_delta" flag:"-segment_time_delta"`
	//Specify the accuracy time when selecting the start time for a segment, expressed as a duration specification. Default value is "0".
	//When delta is specified a key-frame will start a new segment if its PTS satisfies the relation:

	//PTS >= start_time - time_delta
	//This option is useful when splitting video content, which is always split at GOP boundaries, in case a key frame is found just before the specified split time.

	//In particular may be used in combination with the ffmpeg option force_key_frames. The key frame times specified by force_key_frames may not be set accurately because of rounding issues, with the consequence that a key frame time may result set just before the specified time. For constant frame rate videos a value of 1/(2*frame_rate) should address the worst case mismatch between the specified time and the time set by force_key_frames.

	SegmentTimes typex.Time `json:"segment_times" flag:"-segment_times"`
	//Specify a list of split points. times contains a list of comma separated duration specifications, in increasing order. See also the segment_time option.

	SegmentFrames typex.Frames `json:"segment_frames" flag:"-segment_frames"`
	//Specify a list of split video frame numbers. frames contains a list of comma separated integer numbers, in increasing order.

	//This option specifies to start a new segment whenever a reference stream key frame is found and the sequential number (starting from 0) of the frame is greater or equal to the next value in the list.

	SegmentWrap typex.LimitSize `json:"segment_wrap" flag:"-segment_wrap"`
	//Wrap around segment index once it reaches limit.

	SegmentStartNumber typex.UNumber `json:"segment_start_number" flag:"-segment_start_number"`
	//Set the sequence number of the first segment. Defaults to 0.

	Strftime typex.Bool `json:"strftime" flag:"-strftime"`
	//Use the strftime function to define the name of the new segments to write. If this is selected, the output segment name must contain a strftime function template. Default value is 0.

	BreakNonKeyframes typex.Bool `json:"break_non_keyframes" flag:"-break_non_keyframes"`
	//If enabled, allow segments to start on frames other than keyframes. This improves behavior on some players when the time between keyframes is inconsistent, but may make things worse on others, and can cause some oddities during seeking. Defaults to 0.

	ResetTimestamps typex.Bool `json:"reset_timestamps" flag:"-reset_timestamps"`
	//Reset timestamps at the beginning of each segment, so that each segment will start with near-zero timestamps. It is meant to ease the playback of the generated segments. May not work with some combinations of muxers/codecs. It is set to 0 by default.

	InitialOffset typex.U16Offset `json:"initial_offset" flag:"-initial_offset"`
	//Specify timestamp offset to apply to the output packet timestamps. The argument must be a time duration specification, and defaults to 0.

	WriteEmptySegments typex.UBool `json:"write_empty_segments" flag:"-write_empty_segments"`
	//If enabled, write an empty segment if there are no packets during the period a segment would usually span. Otherwise, the segment will be filled with the next packet written. Defaults to 0.

}
