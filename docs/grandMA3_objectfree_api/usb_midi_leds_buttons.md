# USB / MIDI device LEDs & buttons

Drive LEDs on attached USB / MIDI surfaces and read their button state - the core of the APC mini bridge.

## Functions

- [`SetLED`](#setled)
- [`GetButton`](#getbutton)

---


## SetLED

**Signature**

```
SetLED(light_userdata:usb_device_object_handle, table:led_values): nothing
```

**Help page title:** `SetLED(handle,table)`

**Description**

The **SetLED** Lua function sends a table with a set of LED brightness values to an MA3Module. After around two seconds, the system automatically sets the LED values to what it believes it should be.

Below the example is a table listing all the grandMA3 hardware modules and which index number matches which LED on the hardware module.

            **Important:**

Setting a value above 0 for a table index number not connected to an LED can cause the module to crash and reboot

**Arguments**

- **Handle:**

 This function does not accept any arguments.

- **Table:**

 The table should be an indexed table with a set of integer values. The value range is from 0 to 255. This range indicates a brightness level. A special value of "-1" is used to release the LED to the system. The table should contain 1024 indexes.

**Return**

This function does not return anything.

**Hardware Modules LED Table**

This table provides a list of index numbers (indexed from 1). The numbers are matched with elements on the three different kinds of hardware modules that are relevant. Notice that the naming here comes from the internal hardware definition and might not match exactly the print on the keys or the official name.

 Index
 grandMA3 Master Module(MM)
 grandMA3 Fader Module Encoder(MFE)
 grandMA3 Fader Module Crossfader(MFX)

0

1

ENCODER_INSIDE4 Red

Executor 108 Button

Executor 108 Button

2

ENCODER_OUTSIDE3 Red

Executor 110 Button

Executor 110 Button

3

ENCODER_INSIDE2 Red

Executor 211 Button

Executor 211 Button

4

EXEC_GrandKnob Red

Executor 212 Button

Executor 212 Button

5

MENU

Executor 213 Button

Executor 213 Button

6

ENCODER_OUTSIDE2 Red

Executor 214 Button

Executor 214 Button

7

ENCODER_INSIDE1 Red

Executor 215 Button

Executor 215 Button

8

ENCODER_OUTSIDE1 Red

Executor 209 Button

XFade1Btn Knob Red

9

EXEC_GrandKnob  Green

Executor 210 Button

XFade2Btn Knob Red

10

ENCODER_INSIDE1 Green

Executor 208 Button

Executor 209 Button

11

ENCODER_OUTSIDE1 Green

Executor 115 Button

Executor 210 Button

12

ENCODER_INSIDE2 Green

Executor 114 Button

Executor 208 Button

13

ENCODER_OUTSIDE2 Green

Executor 113 Button

XFade1Btn Knob Green

14

ENCODER_INSIDE3 Green

Executor 112 Button

XFade2Btn Knob Green

15

ENCODER_OUTSIDE3 Green

Executor 207 Button

DEF_GO

16

ENCODER_INSIDE4 Green

Executor 206 Button

Executor 115 Button

17

ENCODER_OUTSIDE4 Green

Executor 205 Button

Executor 114 Button

18

ENCODER_INSIDE5 Green

Executor 105 Button

Executor 113 Button

19

ENCODER_OUTSIDE5 Green

Executor 106 Button

Executor 112 Button

20

ENCODER_OUTSIDE4 Red

Executor 107 Button

Executor 207 Button

21

EXEC_GrandKnob Blue

Executor 109 Button

Executor 206 Button

22

ENCODER_INSIDE1 Blue

Executor 204 Button

Executor 205 Button

23

ENCODER_OUTSIDE1 Blue

Executor 203 Button

XFade2 Fader Red

24

ENCODER_INSIDE2 Blue

Executor 202 Button

XFade2 Fader Green

25

ENCODER_OUTSIDE2 Blue

Executor 201 Button

XFade2 Fader Blue

26

ENCODER_INSIDE3 Blue

Executor 111 Button

Executor 105 Button

27

ENCODER_OUTSIDE3 Blue

Executor 101 Button

Executor 106 Button

28

ENCODER_INSIDE4 Blue

Executor 102 Button

Executor 107 Button

29

ENCODER_OUTSIDE4 Blue

Executor 103 Button

XFade1Btn Knob Blue

30

ENCODER_INSIDE5 Blue

Executor 104 Button

XFade2Btn Knob Blue

31

ENCODER_OUTSIDE5 Blue

Executor 312 Fader Red

Executor 109 Button

32

ENCODER_OUTSIDE5 Red

Executor 311 Fader Red

Executor 204 Button

33

ENCODER_INSIDE5 Red

Executor 413 Fader Red

Executor 203 Button

34

ENCODER_INSIDE3 Red

Executor 411 Button

Executor 202 Button

35

ESC

Executor 412 Button

Executor 201 Button

36

CLEAR

Executor 414 Fader Red

XFade2Btn Button

37

HELP

Executor 411 Fader Red

DEF_PAUSE

38

GOTO

Executor 412 Fader Red

XFade1Btn Button

39

ALIGN

Executor 415 Button

Executor 111 Button

40

COPY

Executor 414 Button

Executor 101 Button

41

OFF

Executor 413 Button

Executor 102 Button

42

FULL

Executor 411 Fader Green

Executor 103 Button

43

PLEASE

Executor 412 Fader Green

Executor 104 Button

44

NUM4

Executor 413 Fader Green

DEF_GOBACK

45

NUM5

Executor 414 Fader Green

XFade1 Fader Red

46

UNDO

Executor 415 Fader Green

XFade1 Fader Green

47

GROUP

Executor 311 Fader Green

XFade1 Fader Blue

48

THRU

Executor 312 Fader Green

Executor 312 Fader Red

49

NUM6

Executor 313 Fader Green

Executor 311 Fader Red

50

NUM2

Executor 314 Fader Green

Executor 413 Fader Red

51

STORE

Executor 315 Fader Green

Executor 411 Button

52

ASSIGN

Executor 312 Button

Executor 412 Button

53

AT

Executor 313 Fader Red

Executor 414 Fader Red

54

MA1

Executor 311 Button

Executor 411 Fader Red

55

SLASH

Executor 411 Fader Blue

Executor 412 Fader Red

56

NUM1

Executor 412 Fader Blue

Executor 415 Button

57

CUE

Executor 413 Fader Blue

Executor 414 Button

58

TIME

Executor 414 Fader Blue

Executor 413 Button

59

SEQUENCE

Executor 415 Fader Blue

Executor 411 Fader Green

60

CHANNEL

Executor 311 Fader Blue

Executor 412 Fader Green

61

NUM7

Executor 312 Fader Blue

Executor 413 Fader Green

62

NUM8

Executor 313 Fader Blue

Executor 414 Fader Green

63

NUM9

Executor 314 Fader Blue

Executor 415 Fader Green

64

NUM3

Executor 315 Fader Blue

Executor 311 Fader Green

65

MINUS

Executor 313 Button

Executor 312 Fader Green

66

NUM0

Executor 315 Fader Red

Executor 313 Fader Green

67

DOT

Executor 314 Fader Red

Executor 314 Fader Green

68

IF

Executor 314 Button

Executor 315 Fader Green

69

PLUS

Executor 315 Button

Executor 312 Button

70

LEARN

Executor 415 Fader Red

Executor 313 Fader Red

71

Executor 297 Knob Red

Executor 307 Fader Red

Executor 311 Button

72

Executor 295 "X5 | Step"

Executor 306 Fader Red

Executor 411 Fader Blue

73

Executor 296 "X6 | TC"

Executor 408 Fader Red

Executor 412 Fader Blue

74

Executor 298 Knob Red

Executor 406 Button

Executor 413 Fader Blue

75

Executor 295 Knob Red

Executor 407 Button

Executor 414 Fader Blue

76

Executor 296 Knob Red

Executor 409 Fader Red

Executor 415 Fader Blue

77

GOFAST

Executor 406 Fader Red

Executor 311 Fader Blue

78

Executor 195 "X13 | Phaser"

Executor 407 Fader Red

Executor 312 Fader Blue

79

Executor 196 "X14 | Macro"

Executor 410 Button

Executor 313 Fader Blue

80

Executor 295 Knob Green

Executor 409 Button

Executor 314 Fader Blue

81

Executor 296 Knob Green

Executor 408 Button

Executor 315 Fader Blue

82

Executor 297 Knob Green

Executor 406 Fader Green

Executor 313 Button

83

Executor 298 Knob Green

Executor 407 Fader Green

Executor 315 Fader Red

84

DELETE

Executor 408 Fader Green

Executor 314 Fader Red

85

Executor 197 "X15 | Page"

Executor 409 Fader Green

Executor 314 Button

86

GOBACKFAST

Executor 410 Fader Green

Executor 315 Button

87

Executor 298 "X8 | DMX"

Executor 306 Fader Green

Executor 415 Fader Red

88

STOMP

Executor 307 Fader Green

Executor 307 Fader

89

SELECT

Executor 308 Fader Green

Executor 306 Fader Red

90

Executor 295 Knob Blue

Executor 309 Fader Green

Executor 408 Fader Red

91

Executor 296 Knob Blue

Executor 310 Fader Green

Executor 406 Button

92

Executor 297 Knob Blue

Executor 307 Button

Executor 407 Button

93

Executor 298 Knob Blue

Executor 308 Fader Red

Executor 409 Fader Red

94

Executor 198 "X16 | Exec"

Executor 306 Button

Executor 406 Fader Red

95

Executor 297 "X7 | View"

Executor 406 Fader Blue

Executor 407 Fader Red

96

ON

Executor 407 Fader Blue

Executor 410 Button

97

MOVE

Executor 408 Fader Blue

Executor 409 Button

98

FIXTURE

Executor 409 Fader Blue

Executor 408 Button

99

PRESET

Executor 410 Fader Blue

Executor 406 Fader Green

100

EDIT

Executor 306 Fader Blue

Executor 407 Fader Green

101

UPDATE

Executor 307 Fader Blue

Executor 408 Fader Green

102

PAUSE

Executor 308 Fader Blue

Executor 409 Fader Green

103

GOBACK

Executor 309 Fader Blue

Executor 410 Fader Green

104

Executor 293 Knob Red

Executor 310 Fader Blue

Executor 306 Fader Green

105

SOLO

Executor 308 Button

Executor 307 Fader Green

106

HIGHLIGHT

Executor 310 Fader Red

Executor 308 Fader Green

107

Executor 294 Knob Red

Executor 309 Fader Red

Executor 309 Fader Green

108

Executor 291 Knob Red

Executor 309 Button

Executor 310 Fader Green

109

Executor 292 Knob Red

Executor 310 Button

Executor 307 Button

110

GO

Executor 410 Fader Red

Executor 308 Fader Red

111

LIST

Executor 302 Fader Red

Executor 306 Button

112

PAGE_DOWN

Executor 301 Fader Red

Executor 406 Fader Blue

113

Executor 291 Knob Green

Executor 403 Fader Red

Executor 407 Fader Blue

114

Executor 292 Knob Green

Executor 401 Button

Executor 408 Fader Blue

115

Executor 293 Knob Green

Executor 402 Button

Executor 409 Fader Blue

116

Executor 294 Knob Green

Executor 404 Fader Red

Executor 410 Fader Blue

117

SELFIX

Executor 401 Fader Red

Executor 306 Fader Blue

118

MA2

Executor 402 Fader Red

Executor 307 Fader Blue

119

PAGE_UP

Executor 405 Button

Executor 308 Fader Blue

120

XKEYS

Executor 404 Button

Executor 309 Fader Blue

121

BLIND

Executor 403 Button

Executor 310 Fader Blue

122

Executor 192 "X10"

Executor 401 Fader Green

Executor 308 Button

123

Executor 193 "X11"

Executor 402 Fader Green

Executor 310 Fader Red

124

Executor 194 "X12"

Executor 403 Fader Green

Executor 309 Fader Red

125

Executor 291 Knob Blue

Executor 404 Fader Green

Executor 309 Button

126

Executor 292 Knob Blue

Executor 405 Fader Green

Executor 310 Button

127

Executor 293 Knob Blue

Executor 301 Fader Green

Executor 410 Fader Red

128

Executor 294 Knob Blue

Executor 302 Fader Green

Executor 302 Fader Red

129

PREVIEW

Executor 303 Fader Green

Executor 301 Fader Red

130

FREEZE

Executor 304 Fader Green

Executor 403 Fader Red

131

DOWN

Executor 305 Fader Green

Executor 401 Button

132

PREV

Executor 302 Button

Executor 402 Button

133

RESET

Executor 303 Fader Red

Executor 404 Fader Red

134

UP

Executor 301 Button

Executor 401 Fader Red

135

Executor 291 "X1 | Clone"

Executor 401 Fader Blue

Executor 402 Fader Red

136

Executor 292 "X2 | Link"

Executor 402 Fader Blue

Executor 405 Button

137

Executor 293 "X3 | Grid"

Executor 403 Fader Blue

Executor 404 Button

138

Executor 294 "X4 | Layout"

Executor 404 Fader Blue

Executor 403 Button

139

Executor 191 "X9"

Executor 405 Fader Blue

Executor 401 Fader Green

140

NEXT

Executor 301 Fader Blue

Executor 402 Fader Green

141

All LEDs on the Keyboard

Executor 302 Fader Blue

Executor 403 Fader Green

142

Small Screen Backlight

Executor 303 Fader Blue

Executor 404 Fader Green

143

Letterbox Screen Backlight

Executor 304 Fader Blue

Executor 405 Fader Green

144

Executor 305 Fader Blue

Executor 301 Fader Green

145

Executor 303 Button

Executor 302 Fader Green

146

Executor 305 Fader Red

Executor 303 Fader Green

147

Executor 304 Fader Red

Executor 304 Fader Green

148

Executor 304 Button

Executor 305 Fader Green

149

Executor 305 Button

Executor 302 Button

150

Executor 405 Fader Red

Executor 303 Fader Red

151

RateBtn2

Executor 301 Button

152

ExecBtn1

Executor 401 Fader Blue

153

SpeedBtn1

Executor 402 Fader Blue

154

RateBtn1

Executor 403 Fader Blue

155

SpeedBtn2

Executor 404 Fader Blue

156

ProgBtn1

Executor 405 Fader Blue

157

ProgBtn2

Executor 301 Fader Blue

158

ProgBtn3

Executor 302 Fader Blue

159

ExecBtn3

Executor 303 Fader Blue

160

ExecBtn2

Executor 304 Fader Blue

161

Executor 201 Fader Red

Executor 305 Fader Blue

162

Executor 201 Fader Green

Executor 303 Button

163

Executor 201 Fader Blue

Executor 305 Fader Red

164

Executor 202 Fader Red

Executor 304 Fader Red

165

Executor 202 Fader Green

Executor 304 Button

166

Executor 202 Fader Blue

Executor 305 Button

167

Executor 203 Fader Red

Executor 405 Fader Red

168

Executor 203 Fader Green

Executor 201 Fader Red

169

Executor 203 Fader Blue

Executor 201 Fader Green

170

Executor 204 Fader Red

Executor 201 Fader Blue

171

Executor 204 Fader Green

Executor 202 Fader Red

172

Executor 204 Fader Blue

Executor 202 Fader Green

173

Executor 205 Fader Red

Executor 202 Fader Blue

174

Executor 205 Fader Green

Executor 203 Fader Red

175

Executor 205 Fader Blue

Executor 203 Fader Green

176

Executor 206 Fader Red

Executor 203 Fader Blue

177

Executor 206 Fader Green

Executor 204 Fader Red

178

Executor 206 Fader Blue

Executor 204 Fader Green

179

Executor 207 Fader Red

Executor 204 Fader Blue

180

Executor 207 Fader Green

Executor 205 Fader Red

181

Executor 207 Fader Blue

Executor 205 Fader Green

182

Executor 208 Fader Red

Executor 205 Fader Blue

183

Executor 208 Fader Green

Executor 206 Fader Red

184

Executor 208 Fader Blue

Executor 206 Fader Green

185

Executor 209 Fader Red

Executor 206 Fader Blue

186

Executor 209 Fader Green

Executor 207 Fader Red

187

Executor 209 Fader Blue

Executor 207 Fader Green

188

Executor 210 Fader Red

Executor 207 Fader Blue

189

Executor 210 Fader Green

Executor 208 Fader Red

190

Executor 210 Fader Blue

Executor 208 Fader Green

191

Executor 211 Fader Red

Executor 208 Fader Blue

192

Executor 211 Fader Green

Executor 209 Fader Red

193

Executor 211 Fader Blue

Executor 209 Fader Green

194

Executor 212 Fader Red

Executor 209 Fader Blue

195

Executor 212 Fader Green

Executor 210 Fader Red

196

Executor 212 Fader Blue

Executor 210 Fader Green

197

Executor 213 Fader Red

Executor 210 Fader Blue

198

Executor 213 Fader Green

Executor 211 Fader Red

199

Executor 213 Fader Blue

Executor 211 Fader Green

200

Executor 214 Fader Red

Executor 211 Fader Blue

201

Executor 214 Fader Green

Executor 212 Fader Red

202

Executor 214 Fader Blue

Executor 212 Fader Green

203

Executor 215 Fader Red

Executor 212 Fader Blue

204

Executor 215 Fader Green

Executor 213 Fader Red

205

Executor 215 Fader Blue

Executor 213 Fader Green

206

Desklights

Executor 213 Fader Blue

207

Letterbox Screen Backlight

Executor 214 Fader Red

208

Small Screen Backlight

Executor 214 Fader Green

209

Executor 214 Fader Blue

210

Executor 215 Fader Red

211

Executor 215 Fader Green

212

Executor 215 Fader Blue

213

Desklights

214

Letterbox Screen Backlight

215

Small Screen Backlight

**Example**

This example sets the LEDs on encoder 1 to green on a full-size console:

```lua
return function()
    -- Create the LED table
    local myLedTable = {}
    -- Fill the table with default "release" value
    for index=1,256 do
        myLedTable[index] = -1;
    end
    -- Set values in the table
    -- Encoder_inside1 = green
    myLedTable[7] = 0
    myLedTable[10] = 255
    myLedTable[22] = 0
    -- Encoder_outside1 = green
    myLedTable[8] = 0
    myLedTable[11] = 255
    myLedTable[23] = 0
    -- Get the handle for the MasterModule on a console
    local usbDeviceHandle = Root().UsbNotifier.MA3Modules["UsbDeviceMA3 2"]
    -- Set the values for the LEDs
    SetLED(usbDeviceHandle, myLedTable)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_setled.html>


## GetButton

**Signature**

```
GetButton(light_userdata:usb_device_object_handle): table of boolean:state
```

**Help page title:** `GetButton(handle)`

**Description**

The **GetButton **Lua function returns a key-value pairs table indicating, with a boolean value, whether a button is pressed on an MA3Module.

          Below the example is a table listing all the grandMA3 hardware modules and which index number matches which button on the hardware module.

**Arguments**

- **Handle:**

            The handle for the MA3 module.

**Return**

- **Table**:

            The returned table is a key-value pairs table with a set of 512 pairs with a boolean value. A **true** boolean value indicates that the button is pressed or the fader is touched. The table key is 1-indexed.

**Hardware Modules Button Table**

This table provides a list of index numbers (0-indexed and 1-indexed). The numbers are matched with elements on the three different kinds of hardware modules. Notice that the naming here comes from the internal hardware definition and might not exactly match the print on the keys or the official name.

                0-Index

                1-Index

                  grandMA3

                  Master Module(MM)

                grandMA3

                Fader Module Encoder(MFE)

                grandMA3

                Fader Module Crossfader(MFX)

                0

                1

                1

                2

                2

                3

                ENCODER_INSIDE4

                EXEC_108

                EXEC_108

                3

                4

                ENCODER_OUTSIDE3

                EXEC_110

                EXEC_110

                4

                5

                5

                6

                6

                7

                7

                8

                8

                9

                9

                10

                ENCODER_INSIDE2

                EXEC_211

                EXEC_211

                10

                11

                EXEC_GrandKnob

                EXEC_212

                EXEC_212

                11

                12

                MENU

                EXEC_213

                EXEC_213

                12

                13

                EXEC_214

                EXEC_214

                13

                14

                ENCODER_OUTSIDE2

                EXEC_215

                EXEC_215

                14

                15

                ENCODER_INSIDE1

                15

                16

                ENCODER_OUTSIDE1

                16

                17

                EXEC_209

                EXEC_209

                17

                18

                EXEC_210

                EXEC_210

                18

                19

                EXEC_208

                EXEC_208

                19

                20

                20

                21

                21

                22

                DEF_GO

                22

                23

                EXEC_115

                EXEC_115

                23

                24

                EXEC_114

                EXEC_114

                24

                25

                EXEC_113

                EXEC_113

                25

                26

                EXEC_112

                EXEC_112

                26

                27

                EXEC_207

                EXEC_207

                27

                28

                EXEC_206

                EXEC_206

                28

                29

                EXEC_205

                EXEC_205

                29

                30

                30

                31

                31

                32

                32

                33

                EXEC_105

                EXEC_105

                33

                34

                ENCODER_OUTSIDE4

                EXEC_106

                EXEC_106

                34

                35

                EXEC_107

                EXEC_107

                35

                36

                36

                37

                37

                38

                EXEC_109

                EXEC_109

                38

                39

                EXEC_204

                EXEC_204

                39

                40

                EXEC_203

                EXEC_203

                40

                41

                EXEC_202

                EXEC_202

                41

                42

                EXEC_201

                EXEC_201

                42

                43

                EXEC_XFade2Btn

                43

                44

                DEF_PAUSE

                44

                45

                EXEC_XFade1Btn

                45

                46

                EXEC_111

                EXEC_111

                46

                47

                ENCODER_OUTSIDE5

                EXEC_101

                EXEC_101

                47

                48

                ENCODER_INSIDE5

                EXEC_102

                EXEC_102

                48

                49

                EXEC_103

                EXEC_103

                49

                50

                EXEC_104

                EXEC_104

                50

                51

                ENCODER_INSIDE3

                DEF_GOBACK

                51

                52

                52

                53

                53

                54

                54

                55

                55

                56

                56

                57

                57

                58

                58

                59

                FADER_211

                FADER_211

                59

                60

                FADER_212

                FADER_212

                60

                61

                FADER_213

                FADER_213

                61

                62

                FADER_214

                FADER_214

                62

                63

                FADER_215

                FADER_215

                63

                64

                FADER_XFade1

                64

                65

                65

                66

                66

                67

                ESC

                FADER_312 (Disabled)

                FADER_312 (Disabled)

                67

                68

                CLEAR

                FADER_311 (Disabled)

                FADER_311 (Disabled)

                68

                69

                69

                70

                70

                71

                71

                72

                72

                73

                73

                74

                HELP

                FADER_413 (Disabled)

                FADER_413 (Disabled)

                74

                75

                EXEC_411

                EXEC_411

                75

                76

                GOTO

                EXEC_412

                EXEC_412

                76

                77

                77

                78

                ALIGN

                FADER_414 (Disabled)

                FADER_414 (Disabled)

                78

                79

                COPY

                FADER_411 (Disabled)

                FADER_411 (Disabled)

                79

                80

                OFF

                FADER_412 (Disabled)

                FADER_412 (Disabled)

                80

                81

                FULL

                EXEC_415

                EXEC_415

                81

                82

                EXEC_414

                EXEC_414

                82

                83

                PLEASE

                EXEC_413

                EXEC_413

                83

                84

                NUM4

                84

                85

                NUM5

                85

                86

                UNDO

                86

                87

                GROUP

                87

                88

                THRU

                88

                89

                NUM6

                89

                90

                NUM2

                90

                91

                91

                92

                STORE

                92

                93

                ASSIGN

                93

                94

                94

                95

                95

                96

                96

                97

                AT

                EXEC_312

                EXEC_312

                97

                98

                MA1

                FADER_313 (Disabled)

                FADER_313 (Disabled)

                98

                99

                SLASH

                EXEC_311

                EXEC_311

                99

                100

                NUM1

                100

                101

                CUE

                101

                102

                TIME

                102

                103

                103

                104

                SEQUENCE

                104

                105

                105

                106

                CHANNEL

                106

                107

                NUM7

                107

                108

                NUM8

                108

                109

                NUM9

                109

                110

                NUM3

                EXEC_313

                EXEC_313

                110

                111

                MINUS

                FADER_315 (Disabled)

                FADER_315 (Disabled)

                111

                112

                NUM0

                FADER_314 (Disabled)

                FADER_314 (Disabled)

                112

                113

                DOT

                EXEC_314

                EXEC_314

                113

                114

                IF

                EXEC_315

                EXEC_315

                114

                115

                PLUS

                FADER_415 (Disabled)

                FADER_415 (Disabled)

                115

                116

                116

                117

                117

                118

                118

                119

                119

                120

                120

                121

                FADER_209

                FADER_209

                121

                122

                FADER_210

                FADER_210

                122

                123

                FADER_XFade2

                123

                124

                124

                125

                125

                126

                126

                127

                127

                128

                128

                129

                129

                130

                130

                131

                FADER_307 (Disabled)

                FADER_307 (Disabled)

                131

                132

                LEARN

                FADER_306 (Disabled)

                FADER_306 (Disabled)

                132

                133

                133

                134

                134

                135

                135

                136

                136

                137

                137

                138

                FADER_297 (Disabled)

                FADER_408 (Disabled)

                FADER_408 (Disabled)

                138

                139

                X5

                EXEC_406

                EXEC_406

                139

                140

                X6

                EXEC_407

                EXEC_407

                140

                141

                141

                142

                FADER_298 (Disabled)

                FADER_409 (Disabled)

                FADER_409 (Disabled)

                142

                143

                FADER_295 (Disabled)

                FADER_406 (Disabled)

                FADER_406 (Disabled)

                143

                144

                FADER_296 (Disabled)

                FADER_407 (Disabled)

                FADER_407 (Disabled)

                144

                145

                GOFAST

                EXEC_410

                EXEC_410

                145

                146

                X13

                EXEC_409

                EXEC_409

                146

                147

                X14

                EXEC_408

                EXEC_408

                147

                148

                148

                149

                149

                150

                150

                151

                151

                152

                152

                153

                153

                154

                DELETE

                154

                155

                X15

                155

                156

                GOBACKFAST

                156

                157

                X8

                157

                158

                158

                159

                159

                160

                160

                161

                STOMP

                EXEC_307

                EXEC_307

                161

                162

                SELECT

                FADER_308 (Disabled)

                FADER_308 (Disabled)

                162

                163

                EXEC_306

                EXEC_306

                163

                164

                164

                165

                165

                166

                166

                167

                167

                168

                168

                169

                X16

                169

                170

                X7

                170

                171

                171

                172

                ON

                172

                173

                MOVE

                173

                174

                FIXTURE

                EXEC_308

                EXEC_308

                174

                175

                PRESET

                FADER_310 (Disabled)

                FADER_310 (Disabled)

                175

                176

                EDIT

                FADER_309 (Disabled)

                FADER_309 (Disabled)

                176

                177

                UPDATE

                EXEC_309

                EXEC_309

                177

                178

                EXEC_310

                EXEC_310

                178

                179

                FADER_410 (Disabled)

                FADER_410 (Disabled)

                179

                180

                180

                181

                181

                182

                182

                183

                183

                184

                184

                185

                185

                186

                186

                187

                FADER_203

                FADER_203

                187

                188

                FADER_204

                FADER_204

                188

                189

                FADER_205

                FADER_205

                189

                190

                FADER_206

                FADER_206

                190

                191

                FADER_207

                FADER_207

                191

                192

                FADER_208

                FADER_208

                192

                193

                193

                194

                194

                195

                PAUSE

                FADER_302 (Disabled)

                FADER_302 (Disabled)

                195

                196

                GOBACK

                FADER_301 (Disabled)

                FADER_301 (Disabled)

                196

                197

                197

                198

                198

                199

                199

                200

                200

                201

                201

                202

                FADER_293 (Disabled)

                FADER_403 (Disabled)

                FADER_403 (Disabled)

                202

                203

                SOLO

                EXEC_401

                EXEC_401

                203

                204

                HIGHLIGHT

                EXEC_402

                EXEC_402

                204

                205

                205

                206

                FADER_294 (Disabled)

                FADER_404 (Disabled)

                FADER_404 (Disabled)

                206

                207

                FADER_291 (Disabled)

                FADER_401 (Disabled)

                FADER_401 (Disabled)

                207

                208

                FADER_292 (Disabled)

                FADER_402 (Disabled)

                FADER_402 (Disabled)

                208

                209

                GO

                EXEC_405

                EXEC_405

                209

                210

                LIST

                EXEC_404

                EXEC_404

                210

                211

                PAGE_DOWN

                EXEC_403

                EXEC_403

                211

                212

                212

                213

                213

                214

                214

                215

                215

                216

                SELFIX

                216

                217

                MA2

                217

                218

                218

                219

                PAGE_UP

                219

                220

                XKEYS

                220

                221

                BLIND

                221

                222

                222

                223

                223

                224

                224

                225

                X10

                EXEC_302

                EXEC_302

                225

                226

                X11

                FADER_303 (Disabled)

                FADER_303 (Disabled)

                226

                227

                X12

                EXEC_301

                EXEC_301

                227

                228

                228

                229

                229

                230

                230

                231

                231

                232

                PREVIEW

                232

                233

                FREEZE

                233

                234

                DOWN

                234

                235

                PREV

                235

                236

                SET

                236

                237

                UP

                237

                238

                X1

                EXEC_303

                EXEC_303

                238

                239

                X2

                FADER_305 (Disabled)

                FADER_305 (Disabled)

                239

                240

                X3

                FADER_304 (Disabled)

                FADER_304 (Disabled)

                240

                241

                X4

                EXEC_304

                EXEC_304

                241

                242

                X9

                EXEC_305

                EXEC_305

                242

                243

                NEXT

                FADER_405 (Disabled)

                FADER_405 (Disabled)

                243

                244

                244

                245

                245

                246

                246

                247

                247

                248

                248

                249

                FADER_201

                FADER_201

                249

                250

                FADER_202

                FADER_202

                250

                251

                251

                252

                252

                253

                253

                254

                254

                255

                255

                256

                256

                257

                257

                258

                258

                259

                EXEC_RateBtn2

                259

                260

                EXEC_ExecBtn1

                260

                261

                261

                262

                262

                263

                263

                264

                264

                265

                265

                266

                266

                267

                267

                268

                268

                269

                269

                270

                270

                271

                EXEC_ProgEncoder

                271

                272

                EXEC_ExecEncoder

                272

                273

                273

                274

                274

                275

                275

                276

                276

                277

                277

                278

                278

                279

                279

                280

                280

                281

                281

                282

                282

                283

                283

                284

                284

                285

                285

                286

                286

                287

                287

                288

                288

                289

                EXEC_SpeedBtn1

                289

                290

                EXEC_RateBtn1

                290

                291

                EXEC_SpeedBtn2

                291

                292

                292

                293

                293

                294

                294

                295

                295

                296

                296

                297

                297

                298

                298

                299

                EXEC_ProgBtn1

                299

                300

                EXEC_ProgBtn2

                300

                301

                EXEC_ProgBtn3

                301

                302

                302

                303

                303

                304

                304

                305

                EXEC_ExecBtn3

                305

                306

                EXEC_ExecBtn2

                306

                307

                307

                308

                308

                309

                309

                310

**Example**

This example requests the buttons states on the master module on a grandMA3 full-size console:

```lua
return function()
    --- grandMA3 full-size modules are:
    --- Master Module (MM): "UsbDeviceMA3 2"
    --- Fader Module Encoder (MFE): "UsbDeviceMA3 3"
    --- Fader Module Crossfader (MFX): "UsbDeviceMA3 4"

    -- Get a handle to the Master Module on a grandMA3 full-size.
    local usbDeviceHandle = Root().UsbNotifier.MA3Modules["UsbDeviceMA3 2"]
    -- Create a table with the button status.
    local buttonTable = GetButton(usbDeviceHandle)
    -- Check if the table is nil and then print an error.
    if buttonTable == nil then
        ErrPrintf("nil")
        return
    end
    -- If the table is not nil, then print a usefull feedback about pressed buttons. 
    for key,value in pairs(buttonTable) do
        if tostring(value) == "true" then
            Printf("The button with the index " .. key .. " is pressed.")
        end
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getbutton.html>

---

[Back to index](README.md)
