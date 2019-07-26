// +build wasm

// Copyright 2014 Joseph Hager and the TinyGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webgl

import (
	"errors"
	"syscall/js"
)

const (
	// https://developer.mozilla.org/en-US/docs/Web/API/WebGL_API/Constants
	ACTIVE_ATTRIBUTES                            = 0x8B89
	ACTIVE_UNIFORMS                              = 0x8B86
	ALIASED_LINE_WIDTH_RANGE                     = 0x846E
	ALIASED_POINT_SIZE_RANGE                     = 0x846D
	ALPHA                                        = 0x1906
	ALPHA_BITS                                   = 0x0D55
	ALWAYS                                       = 0x0207
	ARRAY_BUFFER                                 = 0x8892
	ARRAY_BUFFER_BINDING                         = 0x8894
	ATTACHED_SHADERS                             = 0x8B85
	ACTIVE_TEXTURE                               = 0x84E0
	BACK                                         = 0x0405
	BLEND                                        = 0x0BE2
	BLEND_COLOR                                  = 0x8005
	BLEND_DST_ALPHA                              = 0x80CA
	BLEND_DST_RGB                                = 0x80C8
	BLEND_EQUATION                               = 0x8009
	BLEND_EQUATION_ALPHA                         = 0x883D
	BLEND_EQUATION_RGB                           = 0x8009
	BLEND_SRC_ALPHA                              = 0x80CB
	BLEND_SRC_RGB                                = 0x80C9
	BLUE_BITS                                    = 0x0D54
	BOOL                                         = 0x8B56
	BOOL_VEC2                                    = 0x8B57
	BOOL_VEC3                                    = 0x8B58
	BOOL_VEC4                                    = 0x8B59
	BROWSER_DEFAULT_WEBGL                        = 0x9244
	BUFFER_SIZE                                  = 0x8764
	BUFFER_USAGE                                 = 0x8765
	BYTE                                         = 0x1400
	CCW                                          = 0x0901
	CLAMP_TO_EDGE                                = 0x812F
	COLOR_ATTACHMENT0                            = 0x8CE0
	COLOR_BUFFER_BIT                             = 0x00004000
	COLOR_CLEAR_VALUE                            = 0x0C22
	COLOR_WRITEMASK                              = 0x0C23
	COMPILE_STATUS                               = 0x8B81
	COMPRESSED_TEXTURE_FORMATS                   = 0x86A3
	CONSTANT_ALPHA                               = 0x8003
	CONSTANT_COLOR                               = 0x8001
	CONTEXT_LOST_WEBGL                           = 0x9242
	CULL_FACE                                    = 0x0B44
	CULL_FACE_MODE                               = 0x0B45
	CURRENT_PROGRAM                              = 0x8B8D
	CURRENT_VERTEX_ATTRIB                        = 0x8626
	CW                                           = 0x0900
	DECR                                         = 0x1E03
	DECR_WRAP                                    = 0x8508
	DELETE_STATUS                                = 0x8B80
	DEPTH_ATTACHMENT                             = 0x8D00
	DEPTH_BITS                                   = 0x0D56
	DEPTH_BUFFER_BIT                             = 0x00000100
	DEPTH_CLEAR_VALUE                            = 0x0B73
	DEPTH_COMPONENT                              = 0x1902
	DEPTH_COMPONENT16                            = 0x81A5
	DEPTH_FUNC                                   = 0x0B74
	DEPTH_RANGE                                  = 0x0B70
	DEPTH_STENCIL                                = 0x84F9
	DEPTH_STENCIL_ATTACHMENT                     = 0x821A
	DEPTH_TEST                                   = 0x0B71
	DEPTH_WRITEMASK                              = 0x0B72
	DITHER                                       = 0x0BD0
	DONT_CARE                                    = 0x1100
	DST_ALPHA                                    = 0x0304
	DST_COLOR                                    = 0x0306
	DYNAMIC_DRAW                                 = 0x88E8
	ELEMENT_ARRAY_BUFFER                         = 0x8893
	ELEMENT_ARRAY_BUFFER_BINDING                 = 0x8895
	EQUAL                                        = 0x0202
	FASTEST                                      = 0x1101
	FLOAT                                        = 0x1406
	FLOAT_MAT2                                   = 0x8B5A
	FLOAT_MAT3                                   = 0x8B5B
	FLOAT_MAT4                                   = 0x8B5C
	FLOAT_VEC2                                   = 0x8B50
	FLOAT_VEC3                                   = 0x8B51
	FLOAT_VEC4                                   = 0x8B52
	FRAGMENT_SHADER                              = 0x8B30
	FRAMEBUFFER                                  = 0x8D40
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           = 0x8CD1
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           = 0x8CD0
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 0x8CD3
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         = 0x8CD2
	FRAMEBUFFER_BINDING                          = 0x8CA6
	FRAMEBUFFER_COMPLETE                         = 0x8CD5
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT            = 0x8CD6
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS            = 0x8CD9
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    = 0x8CD7
	FRAMEBUFFER_UNSUPPORTED                      = 0x8CDD
	FRONT                                        = 0x0404
	FRONT_AND_BACK                               = 0x0408
	FRONT_FACE                                   = 0x0B46
	FUNC_ADD                                     = 0x8006
	FUNC_REVERSE_SUBTRACT                        = 0x800B
	FUNC_SUBTRACT                                = 0x800A
	GENERATE_MIPMAP_HINT                         = 0x8192
	GEQUAL                                       = 0x0206
	GREATER                                      = 0x0204
	GREEN_BITS                                   = 0x0D53
	HIGH_FLOAT                                   = 0x8DF2
	HIGH_INT                                     = 0x8DF5
	IMPLEMENTATION_COLOR_READ_FORMAT             = 0x8B9B
	IMPLEMENTATION_COLOR_READ_TYPE               = 0x8B9A
	INCR                                         = 0x1E02
	INCR_WRAP                                    = 0x8507
	INT                                          = 0x1404
	INT_VEC2                                     = 0x8B53
	INT_VEC3                                     = 0x8B54
	INT_VEC4                                     = 0x8B55
	INVALID_ENUM                                 = 0x0500
	INVALID_FRAMEBUFFER_OPERATION                = 0x0506
	INVALID_OPERATION                            = 0x0502
	INVALID_VALUE                                = 0x0501
	INVERT                                       = 0x150A
	KEEP                                         = 0x1E00
	LEQUAL                                       = 0x0203
	LESS                                         = 0x0201
	LINEAR                                       = 0x2601
	LINEAR_MIPMAP_LINEAR                         = 0x2703
	LINEAR_MIPMAP_NEAREST                        = 0x2701
	LINES                                        = 0x0001
	LINE_LOOP                                    = 0x0002
	LINE_STRIP                                   = 0x0003
	LINE_WIDTH                                   = 0x0B21
	LINK_STATUS                                  = 0x8B82
	LOW_FLOAT                                    = 0x8DF0
	LOW_INT                                      = 0x8DF3
	LUMINANCE                                    = 0x1909
	LUMINANCE_ALPHA                              = 0x190A
	MAX_COMBINED_TEXTURE_IMAGE_UNITS             = 0x8B4D
	MAX_CUBE_MAP_TEXTURE_SIZE                    = 0x851C
	MAX_FRAGMENT_UNIFORM_VECTORS                 = 0x8DFD
	MAX_RENDERBUFFER_SIZE                        = 0x84E8
	MAX_TEXTURE_IMAGE_UNITS                      = 0x8872
	MAX_TEXTURE_SIZE                             = 0x0D33
	MAX_VARYING_VECTORS                          = 0x8DFC
	MAX_VERTEX_ATTRIBS                           = 0x8869
	MAX_VERTEX_TEXTURE_IMAGE_UNITS               = 0x8B4C
	MAX_VERTEX_UNIFORM_VECTORS                   = 0x8DFB
	MAX_VIEWPORT_DIMS                            = 0x0D3A
	MEDIUM_FLOAT                                 = 0x8DF1
	MEDIUM_INT                                   = 0x8DF4
	MIRRORED_REPEAT                              = 0x8370
	NEAREST                                      = 0x2600
	NEAREST_MIPMAP_LINEAR                        = 0x2702
	NEAREST_MIPMAP_NEAREST                       = 0x2700
	NEVER                                        = 0x0200
	NICEST                                       = 0x1102
	NONE                                         = 0
	NOTEQUAL                                     = 0x0205
	NO_ERROR                                     = 0
	ONE                                          = 1
	ONE_MINUS_CONSTANT_ALPHA                     = 0x8004
	ONE_MINUS_CONSTANT_COLOR                     = 0x8002
	ONE_MINUS_DST_ALPHA                          = 0x0305
	ONE_MINUS_DST_COLOR                          = 0x0307
	ONE_MINUS_SRC_ALPHA                          = 0x0303
	ONE_MINUS_SRC_COLOR                          = 0x0301
	OUT_OF_MEMORY                                = 0x0505
	PACK_ALIGNMENT                               = 0x0D05
	POINTS                                       = 0x0000
	POLYGON_OFFSET_FACTOR                        = 0x8038
	POLYGON_OFFSET_FILL                          = 0x8037
	POLYGON_OFFSET_UNITS                         = 0x2A00
	RED_BITS                                     = 0x0D52
	RENDERBUFFER                                 = 0x8D41
	RENDERBUFFER_ALPHA_SIZE                      = 0x8D53
	RENDERBUFFER_BINDING                         = 0x8CA7
	RENDERBUFFER_BLUE_SIZE                       = 0x8D52
	RENDERBUFFER_DEPTH_SIZE                      = 0x8D54
	RENDERBUFFER_GREEN_SIZE                      = 0x8D51
	RENDERBUFFER_HEIGHT                          = 0x8D43
	RENDERBUFFER_INTERNAL_FORMAT                 = 0x8D44
	RENDERBUFFER_RED_SIZE                        = 0x8D50
	RENDERBUFFER_STENCIL_SIZE                    = 0x8D55
	RENDERBUFFER_WIDTH                           = 0x8D42
	RENDERER                                     = 0x1F01
	REPEAT                                       = 0x2901
	REPLACE                                      = 0x1E01
	RGB                                          = 0x1907
	RGB5_A1                                      = 0x8057
	RGB565                                       = 0x8D62
	RGBA                                         = 0x1908
	RGBA4                                        = 0x8056
	SAMPLER_2D                                   = 0x8B5E
	SAMPLER_CUBE                                 = 0x8B60
	SAMPLES                                      = 0x80A9
	SAMPLE_ALPHA_TO_COVERAGE                     = 0x809E
	SAMPLE_BUFFERS                               = 0x80A8
	SAMPLE_COVERAGE                              = 0x80A0
	SAMPLE_COVERAGE_INVERT                       = 0x80AB
	SAMPLE_COVERAGE_VALUE                        = 0x80AA
	SCISSOR_BOX                                  = 0x0C10
	SCISSOR_TEST                                 = 0x0C11
	SHADER_TYPE                                  = 0x8B4F
	SHADING_LANGUAGE_VERSION                     = 0x8B8C
	SHORT                                        = 0x1402
	SRC_ALPHA                                    = 0x0302
	SRC_ALPHA_SATURATE                           = 0x0308
	SRC_COLOR                                    = 0x0300
	STATIC_DRAW                                  = 0x88E4
	STENCIL_ATTACHMENT                           = 0x8D20
	STENCIL_BACK_FAIL                            = 0x8801
	STENCIL_BACK_FUNC                            = 0x8800
	STENCIL_BACK_PASS_DEPTH_FAIL                 = 0x8802
	STENCIL_BACK_PASS_DEPTH_PASS                 = 0x8803
	STENCIL_BACK_REF                             = 0x8CA3
	STENCIL_BACK_VALUE_MASK                      = 0x8CA4
	STENCIL_BACK_WRITEMASK                       = 0x8CA5
	STENCIL_BITS                                 = 0x0D57
	STENCIL_BUFFER_BIT                           = 0x00000400
	STENCIL_CLEAR_VALUE                          = 0x0B91
	STENCIL_FAIL                                 = 0x0B94
	STENCIL_FUNC                                 = 0x0B92
	STENCIL_INDEX8                               = 0x8D48
	STENCIL_PASS_DEPTH_FAIL                      = 0x0B95
	STENCIL_PASS_DEPTH_PASS                      = 0x0B96
	STENCIL_REF                                  = 0x0B97
	STENCIL_TEST                                 = 0x0B90
	STENCIL_VALUE_MASK                           = 0x0B93
	STENCIL_WRITEMASK                            = 0x0B98
	STREAM_DRAW                                  = 0x88E0
	SUBPIXEL_BITS                                = 0x0D50
	TEXTURE                                      = 0x1702
	TEXTURE0                                     = 0x84C0
	TEXTURE1                                     = 0x84C1
	TEXTURE2                                     = 0x84C2
	TEXTURE3                                     = 0x84C3
	TEXTURE4                                     = 0x84C4
	TEXTURE5                                     = 0x84C5
	TEXTURE6                                     = 0x84C6
	TEXTURE7                                     = 0x84C7
	TEXTURE8                                     = 0x84C8
	TEXTURE9                                     = 0x84C9
	TEXTURE10                                    = 0x84CA
	TEXTURE11                                    = 0x84CB
	TEXTURE12                                    = 0x84CC
	TEXTURE13                                    = 0x84CD
	TEXTURE14                                    = 0x84CE
	TEXTURE15                                    = 0x84CF
	TEXTURE16                                    = 0x84D0
	TEXTURE17                                    = 0x84D1
	TEXTURE18                                    = 0x84D2
	TEXTURE19                                    = 0x84D3
	TEXTURE20                                    = 0x84D4
	TEXTURE21                                    = 0x84D5
	TEXTURE22                                    = 0x84D6
	TEXTURE23                                    = 0x84D7
	TEXTURE24                                    = 0x84D8
	TEXTURE25                                    = 0x84D9
	TEXTURE26                                    = 0x84DA
	TEXTURE27                                    = 0x84DB
	TEXTURE28                                    = 0x84DC
	TEXTURE29                                    = 0x84DD
	TEXTURE30                                    = 0x84DE
	TEXTURE31                                    = 0x84DF
	TEXTURE_2D                                   = 0x0DE1
	TEXTURE_BINDING_2D                           = 0x8069
	TEXTURE_BINDING_CUBE_MAP                     = 0x8514
	TEXTURE_CUBE_MAP                             = 0x8513
	TEXTURE_CUBE_MAP_NEGATIVE_X                  = 0x8516
	TEXTURE_CUBE_MAP_NEGATIVE_Y                  = 0x8518
	TEXTURE_CUBE_MAP_NEGATIVE_Z                  = 0x851A
	TEXTURE_CUBE_MAP_POSITIVE_X                  = 0x8515
	TEXTURE_CUBE_MAP_POSITIVE_Y                  = 0x8517
	TEXTURE_CUBE_MAP_POSITIVE_Z                  = 0x8519
	TEXTURE_MAG_FILTER                           = 0x2800
	TEXTURE_MIN_FILTER                           = 0x2801
	TEXTURE_WRAP_S                               = 0x2802
	TEXTURE_WRAP_T                               = 0x2803
	TRIANGLES                                    = 0x0004
	TRIANGLE_FAN                                 = 0x0006
	TRIANGLE_STRIP                               = 0x0005
	UNPACK_ALIGNMENT                             = 0x0CF5
	UNPACK_COLORSPACE_CONVERSION_WEBGL           = 0x9243
	UNPACK_FLIP_Y_WEBGL                          = 0x9240
	UNPACK_PREMULTIPLY_ALPHA_WEBGL               = 0x9241
	UNSIGNED_BYTE                                = 0x1401
	UNSIGNED_INT                                 = 0x1405
	UNSIGNED_SHORT                               = 0x1403
	UNSIGNED_SHORT_4_4_4_4                       = 0x8033
	UNSIGNED_SHORT_5_5_5_1                       = 0x8034
	UNSIGNED_SHORT_5_6_5                         = 0x8363
	VALIDATE_STATUS                              = 0x8B83
	VENDOR                                       = 0x1F00
	VERSION                                      = 0x1F02
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING           = 0x889F
	VERTEX_ATTRIB_ARRAY_ENABLED                  = 0x8622
	VERTEX_ATTRIB_ARRAY_NORMALIZED               = 0x886A
	VERTEX_ATTRIB_ARRAY_POINTER                  = 0x8645
	VERTEX_ATTRIB_ARRAY_SIZE                     = 0x8623
	VERTEX_ATTRIB_ARRAY_STRIDE                   = 0x8624
	VERTEX_ATTRIB_ARRAY_TYPE                     = 0x8625
	VERTEX_SHADER                                = 0x8B31
	VIEWPORT                                     = 0x0BA2
	ZERO                                         = 0
)

type ContextAttributes struct {
	// If Alpha is true, the drawing buffer has an alpha channel for
	// the purposes of performing OpenGL destination alpha operations
	// and compositing with the page.
	Alpha bool

	// If Depth is true, the drawing buffer has a depth buffer of at least 16 bits.
	Depth bool

	// If Stencil is true, the drawing buffer has a stencil buffer of at least 8 bits.
	Stencil bool

	// If Antialias is true and the implementation supports antialiasing
	// the drawing buffer will perform antialiasing using its choice of
	// technique (multisample/supersample) and quality.
	Antialias bool

	// If PremultipliedAlpha is true the page compositor will assume the
	// drawing buffer contains colors with premultiplied alpha.
	// This flag is ignored if the alpha flag is false.
	PremultipliedAlpha bool

	// If the value is true the buffers will not be cleared and will preserve
	// their values until cleared or overwritten by the author.
	PreserveDrawingBuffer bool
}

// Returns a copy of the default WebGL context attributes.
func DefaultAttributes() *ContextAttributes {
	return &ContextAttributes{true, true, false, true, true, false}
}

type Context struct {
	Object js.Value
}

// NewContext takes an HTML5 canvas object and optional context attributes.
// If an error is returned it means you won't have access to WebGL
// functionality.
func NewContext(canvas *js.Value, ca *ContextAttributes) (*Context, error) {
	if ca == nil {
		ca = DefaultAttributes()
	}

	// Info on Context Attributes: https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/getContext
	// (search for "WebGL context attributes" on the page)
	attrStr := "{" +
		"alpha: " + boolStr(ca.Alpha) +
		", depth: " + boolStr(ca.Depth) +
		", stencil: " + boolStr(ca.Stencil) +
		", antialias: " + boolStr(ca.Antialias) +
		", premultipliedAlpha: " + boolStr(ca.PremultipliedAlpha) +
		", preserveDrawingBuffer: " + boolStr(ca.PreserveDrawingBuffer) +
		"}"
	gl := canvas.Call("getContext", "webgl", attrStr)
	if gl == js.Undefined() {
		gl = canvas.Call("getContext", "experimental-webgl", attrStr)
		if gl == js.Undefined() {
			return nil, errors.New("creating a webgl context has failed")
		}
	}
	ctx := new(Context)
	ctx.Object = gl
	return ctx, nil
}

// Returns the context attributes active on the context. These values might
// be different than what was requested on context creation if the
// browser's implementation doesn't support a feature.
func (c *Context) GetContextAttributes() ContextAttributes {
	ca := c.Object.Call("getContextAttributes")
	return ContextAttributes{
		ca.Get("alpha").Bool(),
		ca.Get("depth").Bool(),
		ca.Get("stencil").Bool(),
		ca.Get("antialias").Bool(),
		ca.Get("premultipliedAlpha").Bool(),
		ca.Get("preservedDrawingBuffer").Bool(),
	}
}

// Specifies the active texture unit.
func (c *Context) ActiveTexture(texture int) {
	c.Object.Call("activeTexture", texture)
}

// Attaches a WebGLShader object to a WebGLProgram object.
func (c *Context) AttachShader(program *js.Value, shader *js.Value) {
	c.Object.Call("attachShader", program, shader)
}

// Binds a generic vertex index to a user-defined attribute variable.
func (c *Context) BindAttribLocation(program *js.Value, index int, name string) {
	c.Object.Call("bindAttribLocation", program, index, name)
}

// Associates a buffer with a buffer target.
func (c *Context) BindBuffer(target int, buffer *js.Value) {
	c.Object.Call("bindBuffer", target, buffer)
}

// Associates a WebGLFramebuffer object with the FRAMEBUFFER bind target.
func (c *Context) BindFramebuffer(target int, framebuffer *js.Value) {
	c.Object.Call("bindFramebuffer", target, framebuffer)
}

// Binds a WebGLRenderbuffer object to be used for rendering.
func (c *Context) BindRenderbuffer(target int, renderbuffer *js.Value) {
	c.Object.Call("bindRenderbuffer", target, renderbuffer)
}

// Binds a named texture object to a target.
func (c *Context) BindTexture(target int, texture *js.Value) {
	c.Object.Call("bindTexture", target, texture)
}

// The GL_BLEND_COLOR may be used to calculate the source and destination blending factors.
func (c *Context) BlendColor(r, g, b, a float64) {
	c.Object.Call("blendColor", r, g, b, a)
}

// Sets the equation used to blend RGB and Alpha values of an incoming source
// fragment with a destination values as stored in the fragment's frame buffer.
func (c *Context) BlendEquation(mode int) {
	c.Object.Call("blendEquation", mode)
}

// Controls the blending of an incoming source fragment's R, G, B, and A values
// with a destination R, G, B, and A values as stored in the fragment's WebGLFramebuffer.
func (c *Context) BlendEquationSeparate(modeRGB, modeAlpha int) {
	c.Object.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

// Sets the blending factors used to combine source and destination pixels.
func (c *Context) BlendFunc(sfactor, dfactor int) {
	c.Object.Call("blendFunc", sfactor, dfactor)
}

// Sets the weighting factors that are used by blendEquationSeparate.
func (c *Context) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	c.Object.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// Creates a buffer in memory and initializes it with array data.
// If no array is provided, the contents of the buffer is initialized to 0.
func (c *Context) BufferData(target int, data interface{}, usage int) {
	c.Object.Call("bufferData", target, data, usage)
}

// Used to modify or update some or all of a data store for a bound buffer object.
func (c *Context) BufferSubData(target int, offset int, data interface{}) {
	c.Object.Call("bufferSubData", target, offset, data)
}

// Returns whether the currently bound WebGLFramebuffer is complete.
// If not complete, returns the reason why.
func (c *Context) CheckFramebufferStatus(target int) int {
	return c.Object.Call("checkFramebufferStatus", target).Int()
}

// Sets all pixels in a specific buffer to the same value.
func (c *Context) Clear(flags int) {
	c.Object.Call("clear", flags)
}

// Specifies color values to use by the clear method to clear the color buffer.
func (c *Context) ClearColor(r, g, b, a float32) {
	c.Object.Call("clearColor", r, g, b, a)
}

// Clears the depth buffer to a specific value.
func (c *Context) ClearDepth(depth float64) {
	c.Object.Call("clearDepth", depth)
}

func (c *Context) ClearStencil(s int) {
	c.Object.Call("clearStencil", s)
}

// Lets you set whether individual colors can be written when
// drawing or rendering to a framebuffer.
func (c *Context) ColorMask(r, g, b, a bool) {
	c.Object.Call("colorMask", r, g, b, a)
}

// Compiles the GLSL shader source into binary data used by the WebGLProgram object.
func (c *Context) CompileShader(shader *js.Value) {
	c.Object.Call("compileShader", shader)
}

// Copies a rectangle of pixels from the current WebGLFramebuffer into a texture image.
func (c *Context) CopyTexImage2D(target, level, internal, x, y, w, h, border int) {
	c.Object.Call("copyTexImage2D", target, level, internal, x, y, w, h, border)
}

// Replaces a portion of an existing 2D texture image with data from the current framebuffer.
func (c *Context) CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, w, h int) {
	c.Object.Call("copyTexSubImage2D", target, level, xoffset, yoffset, x, y, w, h)
}

// Creates and initializes a WebGLBuffer.
func (c *Context) CreateBuffer() *js.Value {
	z := c.Object.Call("createBuffer")
	return &z
}

// Creates and initializes a WebGL Array Buffer.
func (c *Context) CreateArrayBuffer() *js.Value {
	z := c.Object.Call("createBuffer", ARRAY_BUFFER)
	return &z
}

// Returns a WebGLFramebuffer object.
func (c *Context) CreateFramebuffer() *js.Value {
	z := c.Object.Call("createFramebuffer")
	return &z
}

// Creates an empty WebGLProgram object to which vector and fragment
// WebGLShader objects can be bound.
func (c *Context) CreateProgram() *js.Value {
	z := c.Object.Call("createProgram")
	return &z
}

// Creates and returns a WebGLRenderbuffer object.
func (c *Context) CreateRenderbuffer() *js.Value {
	z := c.Object.Call("createRenderbuffer")
	return &z
}

// Returns an empty vertex or fragment shader object based on the type specified.
func (c *Context) CreateShader(typ int) *js.Value {
	z := c.Object.Call("createShader", typ)
	return &z
}

// Used to generate a WebGLTexture object to which images can be bound.
func (c *Context) CreateTexture() *js.Value {
	z := c.Object.Call("createTexture")
	return &z
}

// Sets whether or not front, back, or both facing facets are able to be culled.
func (c *Context) CullFace(mode int) {
	c.Object.Call("cullFace", mode)
}

// Delete a specific buffer.
func (c *Context) DeleteBuffer(buffer *js.Value) {
	c.Object.Call("deleteBuffer", buffer)
}

// Deletes a specific WebGLFramebuffer object. If you delete the
// currently bound framebuffer, the default framebuffer will be bound.
// Deleting a framebuffer detaches all of its attachments.
func (c *Context) DeleteFramebuffer(framebuffer *js.Value) {
	c.Object.Call("deleteFramebuffer", framebuffer)
}

// Flags a specific WebGLProgram object for deletion if currently active.
// It will be deleted when it is no longer being used.
// Any shader objects associated with the program will be detached.
// They will be deleted if they were already flagged for deletion.
func (c *Context) DeleteProgram(program *js.Value) {
	c.Object.Call("deleteProgram", program)
}

// Deletes the specified renderbuffer object. If the renderbuffer is
// currently bound, it will become unbound. If the renderbuffer is
// attached to the currently bound framebuffer, it is detached.
func (c *Context) DeleteRenderbuffer(renderbuffer *js.Value) {
	c.Object.Call("deleteRenderbuffer", renderbuffer)
}

// Deletes a specific shader object.
func (c *Context) DeleteShader(shader *js.Value) {
	c.Object.Call("deleteShader", shader)
}

// Deletes a specific texture object.
func (c *Context) DeleteTexture(texture *js.Value) {
	c.Object.Call("deleteTexture", texture)
}

// Sets a function to use to compare incoming pixel depth to the
// current depth buffer value.
func (c *Context) DepthFunc(fun int) {
	c.Object.Call("depthFunc", fun)
}

// Sets whether or not you can write to the depth buffer.
func (c *Context) DepthMask(flag bool) {
	c.Object.Call("depthMask", flag)
}

// Sets the depth range for normalized coordinates to canvas or viewport depth coordinates.
func (c *Context) DepthRange(zNear, zFar float64) {
	c.Object.Call("depthRange", zNear, zFar)
}

// Detach a shader object from a program object.
func (c *Context) DetachShader(program, shader *js.Value) {
	c.Object.Call("detachShader", program, shader)
}

// Turns off specific WebGL capabilities for this context.
func (c *Context) Disable(cap int) {
	c.Object.Call("disable", cap)
}

// Turns off a vertex attribute array at a specific index position.
func (c *Context) DisableVertexAttribArray(index int) {
	c.Object.Call("disableVertexAttribArray", index)
}

// Render geometric primitives from bound and enabled vertex data.
func (c *Context) DrawArrays(mode, first, count int) {
	c.Object.Call("drawArrays", mode, first, count)
}

// Renders geometric primitives indexed by element array data.
func (c *Context) DrawElements(mode, count, typ, offset int) {
	c.Object.Call("drawElements", mode, count, typ, offset)
}

// Turns on specific WebGL capabilities for this context.
func (c *Context) Enable(cap int) {
	c.Object.Call("enable", cap)
}

// Turns on a vertex attribute at a specific index position in
// a vertex attribute array.
func (c *Context) EnableVertexAttribArray(index int) {
	c.Object.Call("enableVertexAttribArray", index)
}

func (c *Context) Finish() {
	c.Object.Call("finish")
}

func (c *Context) Flush() {
	c.Object.Call("flush")
}

// Attaches a WebGLRenderbuffer object as a logical buffer to the
// currently bound WebGLFramebuffer object.
func (c *Context) FrameBufferRenderBuffer(target, attachment, renderbufferTarget int, renderbuffer *js.Value) {
	c.Object.Call("framebufferRenderBuffer", target, attachment, renderbufferTarget, renderbuffer)
}

// Attaches a texture to a WebGLFramebuffer object.
func (c *Context) FramebufferTexture2D(target, attachment, textarget int, texture *js.Value, level int) {
	c.Object.Call("framebufferTexture2D", target, attachment, textarget, texture, level)
}

// Sets whether or not polygons are considered front-facing based
// on their winding direction.
func (c *Context) FrontFace(mode int) {
	c.Object.Call("frontFace", mode)
}

// Creates a set of textures for a WebGLTexture object with image
// dimensions from the original size of the image down to a 1x1 image.
func (c *Context) GenerateMipmap(target int) {
	c.Object.Call("generateMipmap", target)
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a vertex attribute at a specific index position in a program object.
func (c *Context) GetActiveAttrib(program *js.Value, index int) *js.Value {
	z := c.Object.Call("getActiveAttrib", program, index)
	return &z
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a uniform attribute at a specific index position in a program object.
func (c *Context) GetActiveUniform(program *js.Value, index int) *js.Value {
	z := c.Object.Call("getActiveUniform", program, index)
	return &z
}

// Returns a slice of WebGLShaders bound to a WebGLProgram.
func (c *Context) GetAttachedShaders(program *js.Value) []*js.Value {
	objs := c.Object.Call("getAttachedShaders", program)
	shaders := make([]*js.Value, objs.Length())
	for i := 0; i < objs.Length(); i++ {
		z := objs.Index(i)
		shaders[i] = &z
	}
	return shaders
}

// Returns an index to the location in a program of a named attribute variable.
func (c *Context) GetAttribLocation(program *js.Value, name string) int {
	return c.Object.Call("getAttribLocation", program, name).Int()
}

// TODO: Create type specific variations.
// Returns the type of a parameter for a given buffer.
func (c *Context) GetBufferParameter(target, pname int) *js.Value {
	z := c.Object.Call("getBufferParameter", target, pname)
	return &z
}

// TODO: Create type specific variations.
// Returns the natural type value for a constant parameter.
func (c *Context) GetParameter(pname int) *js.Value {
	z := c.Object.Call("getParameter", pname)
	return &z
}

// Returns a value for the WebGL error flag and clears the flag.
func (c *Context) GetError() int {
	return c.Object.Call("getError").Int()
}

// TODO: Create type specific variations.
// Enables a passed extension, otherwise returns null.
func (c *Context) GetExtension(name string) *js.Value {
	z := c.Object.Call("getExtension", name)
	return &z
}

// TODO: Create type specific variations.
// Gets a parameter value for a given target and attachment.
func (c *Context) GetFramebufferAttachmentParameter(target, attachment, pname int) *js.Value {
	z := c.Object.Call("getFramebufferAttachmentParameter", target, attachment, pname)
	return &z
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as an int.
func (c *Context) GetProgramParameteri(program *js.Value, pname int) int {
	return c.Object.Call("getProgramParameter", program, pname).Int()
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as a bool.
func (c *Context) GetProgramParameterb(program *js.Value, pname int) bool {
	return c.Object.Call("getProgramParameter", program, pname).Bool()
}

// Returns information about the last error that occurred during
// the failed linking or validation of a WebGL program object.
func (c *Context) GetProgramInfoLog(program *js.Value) string {
	return c.Object.Call("getProgramInfoLog", program).String()
}

// TODO: Create type specific variations.
// Returns a renderbuffer parameter from the currently bound WebGLRenderbuffer object.
func (c *Context) GetRenderbufferParameter(target, pname int) *js.Value {
	z := c.Object.Call("getRenderbufferParameter", target, pname)
	return &z
}

// TODO: Create type specific variations.
// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameter(shader *js.Value, pname int) *js.Value {
	z := c.Object.Call("getShaderParameter", shader, pname)
	return &z
}

// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameterb(shader *js.Value, pname int) bool {
	return c.Object.Call("getShaderParameter", shader, pname).Bool()
}

// Returns errors which occur when compiling a shader.
func (c *Context) GetShaderInfoLog(shader *js.Value) string {
	return c.Object.Call("getShaderInfoLog", shader).String()
}

// Returns source code string associated with a shader object.
func (c *Context) GetShaderSource(shader *js.Value) string {
	return c.Object.Call("getShaderSource", shader).String()
}

// Returns a slice of supported extension strings.
func (c *Context) GetSupportedExtensions() []string {
	ext := c.Object.Call("getSupportedExtensions")
	extensions := make([]string, ext.Length())
	for i := 0; i < ext.Length(); i++ {
		extensions[i] = ext.Index(i).String()
	}
	return extensions
}

// TODO: Create type specific variations.
// Returns the value for a parameter on an active texture unit.
func (c *Context) GetTexParameter(target, pname int) *js.Value {
	z := c.Object.Call("getTexParameter", target, pname)
	return &z
}

// TODO: Create type specific variations.
// Gets the uniform value for a specific location in a program.
func (c *Context) GetUniform(program, location *js.Value) *js.Value {
	z := c.Object.Call("getUniform", program, location)
	return &z
}

// Returns a WebGLUniformLocation object for the location
// of a uniform variable within a WebGLProgram object.
func (c *Context) GetUniformLocation(program *js.Value, name string) *js.Value {
	z := c.Object.Call("getUniformLocation", program, name)
	return &z
}

// TODO: Create type specific variations.
// Returns data for a particular characteristic of a vertex
// attribute at an index in a vertex attribute array.
func (c *Context) GetVertexAttrib(index, pname int) *js.Value {
	z := c.Object.Call("getVertexAttrib", index, pname)
	return &z
}

// Returns the address of a specified vertex attribute.
func (c *Context) GetVertexAttribOffset(index, pname int) int {
	return c.Object.Call("getVertexAttribOffset", index, pname).Int()
}

// public function hint(target:GLenum, mode:GLenum) : Void;

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsBuffer(buffer *js.Value) bool {
	return c.Object.Call("isBuffer", buffer).Bool()
}

// Returns whether the WebGL context has been lost.
func (c *Context) IsContextLost() bool {
	return c.Object.Call("isContextLost").Bool()
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsFramebuffer(framebuffer *js.Value) bool {
	return c.Object.Call("isFramebuffer", framebuffer).Bool()
}

// Returns true if program object is valid, false otherwise.
func (c *Context) IsProgram(program *js.Value) bool {
	return c.Object.Call("isProgram", program).Bool()
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsRenderbuffer(renderbuffer *js.Value) bool {
	return c.Object.Call("isRenderbuffer", renderbuffer).Bool()
}

// Returns true if shader is valid, false otherwise.
func (c *Context) IsShader(shader *js.Value) bool {
	return c.Object.Call("isShader", shader).Bool()
}

// Returns true if texture is valid, false otherwise.
func (c *Context) IsTexture(texture *js.Value) bool {
	return c.Object.Call("isTexture", texture).Bool()
}

// Returns whether or not a WebGL capability is enabled for this context.
func (c *Context) IsEnabled(capability int) bool {
	return c.Object.Call("isEnabled", capability).Bool()
}

// Sets the width of lines in WebGL.
func (c *Context) LineWidth(width float64) {
	c.Object.Call("lineWidth", width)
}

// Links an attached vertex shader and an attached fragment shader
// to a program so it can be used by the graphics processing unit (GPU).
func (c *Context) LinkProgram(program *js.Value) {
	c.Object.Call("linkProgram", program)
}

// Sets pixel storage modes for readPixels and unpacking of textures
// with texImage2D and texSubImage2D.
func (c *Context) PixelStorei(pname, param int) {
	c.Object.Call("pixelStorei", pname, param)
}

// Sets the implementation-specific units and scale factor
// used to calculate fragment depth values.
func (c *Context) PolygonOffset(factor, units float64) {
	c.Object.Call("polygonOffset", factor, units)
}

// TODO: Figure out if pixels should be a slice.
// Reads pixel data into an ArrayBufferView object from a
// rectangular area in the color buffer of the active frame buffer.
func (c *Context) ReadPixels(x, y, width, height, format, typ int, pixels *js.Value) {
	c.Object.Call("readPixels", x, y, width, height, format, typ, pixels)
}

// Creates or replaces the data store for the currently bound WebGLRenderbuffer object.
func (c *Context) RenderbufferStorage(target, internalFormat, width, height int) {
	c.Object.Call("renderbufferStorage", target, internalFormat, width, height)
}

//func (c *Context) SampleCoverage(value float64, invert bool) {
//	c.Object.Call("sampleCoverage", value, invert)
//}

// Sets the dimensions of the scissor box.
func (c *Context) Scissor(x, y, width, height int) {
	c.Object.Call("scissor", x, y, width, height)
}

// Sets and replaces shader source code in a shader object.
func (c *Context) ShaderSource(shader *js.Value, source string) {
	c.Object.Call("shaderSource", shader, source)
}

// public function stencilFunc(func:GLenum, ref:GLint, mask:GLuint) : Void;
// public function stencilFuncSeparate(face:GLenum, func:GLenum, ref:GLint, mask:GLuint) : Void;
// public function stencilMask(mask:GLuint) : Void;
// public function stencilMaskSeparate(face:GLenum, mask:GLuint) : Void;
// public function stencilOp(fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;
// public function stencilOpSeparate(face:GLenum, fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;

// Loads the supplied pixel data into a texture.
func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, image *js.Value) {
	c.Object.Call("texImage2D", target, level, internalFormat, format, kind, image)
}

// Sets texture parameters for the current texture unit.
func (c *Context) TexParameteri(target int, pname int, param int) {
	c.Object.Call("texParameteri", target, pname, param)
}

// Replaces a portion of an existing 2D texture image with all of another image.
func (c *Context) TexSubImage2D(target, level, xoffset, yoffset, format, typ int, image *js.Value) {
	c.Object.Call("texSubImage2D", target, level, xoffset, yoffset, format, typ, image)
}

// Assigns a floating point value to a uniform variable for the current program object.
func (c *Context) Uniform1f(location *js.Value, x float32) {
	c.Object.Call("uniform1f", location, x)
}

// Assigns a integer value to a uniform variable for the current program object.
func (c *Context) Uniform1i(location *js.Value, x int) {
	c.Object.Call("uniform1i", location, x)
}

// Assigns 2 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform2f(location *js.Value, x, y float32) {
	c.Object.Call("uniform2f", location, x, y)
}

// Assigns 2 integer values to a uniform variable for the current program object.
func (c *Context) Uniform2i(location *js.Value, x, y int) {
	c.Object.Call("uniform2i", location, x, y)
}

// Assigns 3 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform3f(location *js.Value, x, y, z float32) {
	c.Object.Call("uniform3f", location, x, y, z)
}

// Assigns 3 integer values to a uniform variable for the current program object.
func (c *Context) Uniform3i(location *js.Value, x, y, z int) {
	c.Object.Call("uniform3i", location, x, y, z)
}

// Assigns 4 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform4f(location *js.Value, x, y, z, w float32) {
	c.Object.Call("uniform4f", location, x, y, z, w)
}

// Assigns 4 integer values to a uniform variable for the current program object.
func (c *Context) Uniform4i(location *js.Value, x, y, z, w int) {
	c.Object.Call("uniform4i", location, x, y, z, w)
}

// public function uniform1fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform1iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform2fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform2iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform3fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform3iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform4fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform4iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;

// Sets values for a 2x2 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix2fv(location *js.Value, transpose bool, value []float32) {
	c.Object.Call("uniformMatrix2fv", location, transpose, js.TypedArrayOf(value))
}

// Sets values for a 3x3 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix3fv(location *js.Value, transpose bool, value []float32) {
	c.Object.Call("uniformMatrix3fv", location, transpose, js.TypedArrayOf(value))
}

// Sets values for a 4x4 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix4fv(location *js.Value, transpose bool, value []float32) {
	c.Object.Call("uniformMatrix4fv", location, transpose, js.TypedArrayOf(value))
}

// Set the program object to use for rendering.
func (c *Context) UseProgram(program *js.Value) {
	c.Object.Call("useProgram", program)
}

// Returns whether a given program can run in the current WebGL state.
func (c *Context) ValidateProgram(program *js.Value) {
	c.Object.Call("validateProgram", program)
}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	c.Object.Call("vertexAttribPointer", index, size, typ, normal, stride, offset)
}

// public function vertexAttrib1f(indx:GLuint, x:GLfloat) : Void;
// public function vertexAttrib2f(indx:GLuint, x:GLfloat, y:GLfloat) : Void;
// public function vertexAttrib3f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat) : Void;
// public function vertexAttrib4f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat, w:GLfloat) : Void;
// public function vertexAttrib1fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib2fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib3fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib4fv(indx:GLuint, values:ArrayAccess<Float>) : Void;

// Represents a rectangular viewable area that contains
// the rendering results of the drawing buffer.
func (c *Context) Viewport(x, y, width, height int) {
	c.Object.Call("viewport", x, y, width, height)
}

// Returns true or false value as a string
func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
