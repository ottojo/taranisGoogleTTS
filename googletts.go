package main

type AudioEncoding string

type AudioConfig struct {
	AudioEncoding   AudioEncoding `json:"audioEncoding"`
	SpeakingRate    float64       `json:"speakingRate"`
	Pitch           float64       `json:"pitch"`
	VolumeGainDb    float64       `json:"volumeGainDb"`
	SampleRateHertz float64       `json:"sampleRateHertz"`
}

const (
	AUDIO_ENCODING_UNSPECIFIED AudioEncoding = "AUDIO_ENCODING_UNSPECIFIED" // Not specified. Will return result google.rpc.Code.INVALID_ARGUMENT.
	LINEAR16                   AudioEncoding = "LINEAR16"                   // Uncompressed 16-bit signed little-endian samples (Linear PCM). Audio content returned as LINEAR16 also contains a WAV header.
	MP3                        AudioEncoding = "MP3"                        // MP3 audio.
	OGG_OPUS                   AudioEncoding = "OGG_OPUS"                   // Opus encoded audio wrapped in an ogg container. The result will be a file which can be played natively on Android, and in browsers (at least Chrome and Firefox). The quality of the encoding is considerably higher than MP3 while using approximately the same bitrate.
)

type VoiceSelectionParams struct {
	LanguageCode string          `json:"languageCode"`
	Name         string          `json:"name"`
	SsmlGender   SsmlVoiceGender `json:"ssmlGender"`
}

type SsmlVoiceGender string

const (
	SSML_VOICE_GENDER_UNSPECIFIED SsmlVoiceGender = "SSML_VOICE_GENDER_UNSPECIFIED" // An unspecified gender. In VoiceSelectionParams, this means that the client doesn't care which gender the selected voice will have. In the Voice field of ListVoicesResponse, this may mean that the voice doesn't fit any of the other categories in this enum, or that the gender of the voice isn't known.
	MALE                          SsmlVoiceGender = "MALE"                          // A male voice.
	FEMALE                        SsmlVoiceGender = "FEMALE"                        // A female voice.
	NEUTRAL                       SsmlVoiceGender = "NEUTRAL"                       // A gender-neutral voice.
)

type SynthesisInput struct{
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

type Response struct {
	AudioContent string `json:"audioContent"`
}

type Request struct {
	Input       SynthesisInput       `json:"input"`
	Voice       VoiceSelectionParams `json:"voice"`
	AudioConfig AudioConfig          `json:"audioConfig"`
}
