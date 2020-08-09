package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width              = 600
	height             = 600
	vertexShaderSource = `#version 410 core

	in vec3 vertex_position;
	
	void main()
	{
		gl_Position = vec4(vertex_position, 1.0);
	}
` + "\x00"

	fragmentShaderSource = `#version 410

uniform float rect_width;
uniform float rect_height;
uniform vec2 area_w;
uniform vec2 area_h;
uniform uint max_iterations;

out vec4 pixel_color;

const float color_map[48]= float[](
	0.26, 0.12, 0.06,
	0.1, 0.03, 0.1,
	0.04, 0.0, 0.18,
	0.02, 0.02, 0.29,
	0.0, 0.03, 0.39,
	0.05, 0.17, 0.54,
	0.09, 0.32, 0.69,
	0.22, 0.49, 0.82,
	0.53, 0.71, 0.9,
	0.83, 0.93, 0.97,
	0.95, 0.91, 0.75,
	0.97, 0.79, 0.37,
	1.0, 0.67, 0.0,
	0.8, 0.5, 0.0,
	0.6, 0.34, 0.0,
	0.42, 0.2, 0.01
);
void main()
{
    vec2 C = vec2(gl_FragCoord.x * (area_w.y - area_w.x) / rect_width  + area_w.x,
                        gl_FragCoord.y * (area_h.y - area_h.x) / rect_height + area_h.x);
    vec2 Z = vec2(0.0);
    uint iteration = 0;

    while (iteration < max_iterations)
    {
        float x = Z.x * Z.x - Z.y * Z.y + C.x;
        float y = 2.0 * Z.x * Z.y       + C.y;

        if (x * x + y * y > 4.0)
            break;

        Z.x = x;
        Z.y = y;

        ++iteration;
    }

    uint row_index = (iteration * 100 / max_iterations % 16) * 3;
	pixel_color = vec4((iteration == max_iterations ? vec3(0.0) : vec3(color_map[row_index],
						color_map[row_index+1], color_map[row_index+2])), 1.0);
}` + "\x00"
)

// 0.0,  0.0,  0.0,
// 0.26, 0.18, 0.06,
// 0.1,  0.03, 0.1,
// 0.04, 0.0,  0.18,
// 0.02, 0.02, 0.29,
// 0.0,  0.03, 0.39,
// 0.05, 0.17, 0.54,
// 0.09, 0.32, 0.69,
// 0.22, 0.49, 0.82,
// 0.52, 0.71, 0.9,
// 0.82, 0.92, 0.97,
// 0.94, 0.91, 0.75,
// 0.97, 0.79, 0.37,
// 1.0,  0.67, 0.0,
// 0.8,  0.5,  0.0,
// 0.6,  0.34, 0.0,
// 0.41, 0.2,  0.0

var vboArray [width * height * 3]uint8

var (
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}

	rect = []float32{
		-1.0, -1.0, 0.0,
		1.0, -1.0, 0.0,
		1.0, 1.0, 0.0,
		1.0, 1.0, 0.0,
		-1.0, 1.0, 0.0,
		-1.0, -1.0, 0.0,
	}
)

func createRectangleBuffer() uint32 {

	var buff uint32
	gl.GenBuffers(1, &buff)
	gl.BindBuffer(gl.ARRAY_BUFFER, buff)
	gl.BufferData(gl.ARRAY_BUFFER, len(rect)*4, gl.Ptr(rect), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	return buff
}

func createVAO(rectangleBuff uint32) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, rectangleBuff) // error
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return vao
}

type mandelbrot struct {
	scale   float32
	x       float32
	y       float32
	maxIter uint32
}

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	prog := initOpenGL()

	mandelData := mandelbrot{
		scale:   1.0,
		x:       0.0,
		y:       0.0,
		maxIter: 30,
	}
	var name int32
	for !window.ShouldClose() {
		gl.UseProgram(prog)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		name = gl.GetUniformLocation(prog, gl.Str("rect_width\x00"))
		if name < 0 {
			log.Fatalln("Could not find name")
		}
		gl.Uniform1f(name, width)

		name = gl.GetUniformLocation(prog, gl.Str("rect_height\x00"))
		if name < 0 {
			log.Fatalln("Could not find name")
		}
		gl.Uniform1f(name, height)

		name = gl.GetUniformLocation(prog, gl.Str("area_w\x00"))
		if name < 0 {
			log.Fatalln("Could not find name")
		}
		gl.Uniform2f(name, -2.0*mandelData.scale+mandelData.x,
			mandelData.scale+mandelData.x)

		name = gl.GetUniformLocation(prog, gl.Str("area_h\x00"))
		if name < 0 {
			log.Fatalln("Could not find name")
		}
		gl.Uniform2f(name, -1*mandelData.scale+mandelData.y,
			mandelData.scale+mandelData.y)

		name = gl.GetUniformLocation(prog, gl.Str("max_iterations\x00"))
		if name < 0 {
			log.Fatalln("Could not find name")
		}
		gl.Uniform1ui(name, mandelData.maxIter)

		rectBuf := createRectangleBuffer()
		vao := createVAO(rectBuf)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
		gl.BindVertexArray(0)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Fractals", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
