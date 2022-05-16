package emulator

// Adapted from https://github.com/veandco/go-sdl2-examples/blob/master/examples/playing-audio/playing-audio.go

// typedef unsigned char Uint8;
// void SineWave(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"math"
	"reflect"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	DefaultFrequency = 5000
	DefaultFormat    = sdl.AUDIO_S16
	DefaultChannels  = 2
	DefaultSamples   = 512

	toneHz = 440
	dPhase = 2 * math.Pi * toneHz / DefaultSamples
)

//export SineWave
func SineWave(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length) / 2
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]C.ushort)(unsafe.Pointer(&hdr))

	var phase float64
	for i := 0; i < n; i++ {
		phase += dPhase
		sample := C.ushort((math.Sin(phase) + 0.999999) * 32768)
		buf[i] = sample
	}
}

type speaker struct {
	device sdl.AudioDeviceID
}

func NewSpeaker() (speaker, func()) {
	var dev sdl.AudioDeviceID
	var err error

	spec := sdl.AudioSpec{
		Freq:     DefaultFrequency,
		Format:   DefaultFormat,
		Channels: DefaultChannels,
		Samples:  DefaultSamples,
		Callback: sdl.AudioCallback(C.SineWave),
	}

	if dev, err = sdl.OpenAudioDevice("", false, &spec, nil, 0); err != nil {
		panic("Cannot initialize audio device")
	}

	return speaker{device: dev}, func() { sdl.CloseAudioDevice(dev) }
}

func (s *speaker) beep(play bool) {
	sdl.PauseAudioDevice(s.device, !play)
}
