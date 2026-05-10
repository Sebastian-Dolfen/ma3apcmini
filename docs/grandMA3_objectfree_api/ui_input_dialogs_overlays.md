# UI: input, dialogs & overlays

Show modal dialogs (TextInput / Confirm / MessageBox), inspect displays and focus, and manipulate overlays.

## Functions

- [`TextInput`](#textinput)
- [`PopupInput`](#popupinput)
- [`Confirm`](#confirm)
- [`MessageBox`](#messagebox)
- [`GetDisplayByIndex`](#getdisplaybyindex)
- [`GetRemoteVideoInfo`](#getremotevideoinfo)
- [`GetUIObjectAtPosition`](#getuiobjectatposition)
- [`DrawPointer`](#drawpointer)
- [`WaitObjectDelete`](#waitobjectdelete)
- [`GetFocus`](#getfocus)
- [`GetFocusDisplay`](#getfocusdisplay)
- [`GetDisplayCollect`](#getdisplaycollect)
- [`FindBestFocus`](#findbestfocus)
- [`FindNextFocus`](#findnextfocus)
- [`CloseAllOverlays`](#closealloverlays)
- [`GetTopModal`](#gettopmodal)
- [`GetTopOverlay`](#gettopoverlay)
- [`WaitModal`](#waitmodal)
- [`SetBlockInput`](#setblockinput)
- [`GetBlockInput`](#getblockinput)
- [`FindTexture`](#findtexture)
- [`GetScreenContent`](#getscreencontent)
- [`FSExtendedModeHasDots`](#fsextendedmodehasdots)

---


## TextInput

**Signature**

```
TextInput([string:title[, string:value[, integer:x[, integer:y]]]]): string:value
```

**Help page title:** `TextInput([string[, string[, integer[, integer]]]])`

**Description**

The **TextInput** Lua function opens a text input pop-up and returns the typed input as a string. It is part of the user interface functions.

**Arguments**

- **String** (optional):

            This string is the title for the pop-up. The title bar has a default "Edit" text at the beginning of the title that cannot be removed.

- **String **(optional):

            This string is the text already in the input field - can be used to provide user guidance.

- **Integer** (optional):

            This integer defines a position on the x-axis where the pop-up should appear (on all screens). "0" is on the left side of the screen. Nil or undefined is centered.

- **Integer** (optional):

            This integer defines a position on the y-axis where the pop-up should appear (on all screens). "0" is at the top of the screen. Nil or undefined is centered.

**Return**

- **String**:

            The returned user input.

          Example

          To open a text input and print the entered value in the Command Line History, create a plugin with this code:

```lua
return function()
    -- Create a pop-up with the title and an input field containing some default text
    -- The returned text is store in a Lua variable
    local input = TextInput("This is the title","Please provide your input here")
    -- Print the returned text value
    Printf("You entered this message: %s",tostring(input))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_textinput.html>


## PopupInput

**Signature**

```
PopupInput({title:str, caller:handle, items:table:{{'str'|'int'|'lua'|'handle', name, type-dependent}...}, selectedValue:str, x:int, y:int, target:handle, render_options:{left_icon, number, right_icon}, useTopLeft:bool, properties:{prop:value}, add_args:{FilterSupport='Yes'/'No'}}): integer:selected_index, string:selected_value
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## Confirm

**Signature**

```
Confirm([string:title [,string:message [,integer:display_index [,boolean:showCancel]]]]): boolean:result
```

**Help page title:** `Confirm(string[, string[, integer[, boolean]]])`

**Description**

The **Confirm** Lua function provides a simple confirmation pop-up for a true/false query. It is part of the user interface functions.

**Arguments**

- **String**:

 This string is the title for the pop-up.

- **String **(optional):

 This string is the text in the pop-up.

- **Integer** (optional):

 This integer is not used since the pop-up appears on all screens. The value can be *nil*.

- **Boolean **(optional):

 This boolean defines if there is a Cancel button in the pop-up or not.

  - true: There is a Cancel button in the pop-up. This is the default option used if it is not defined.

  - false: There is only an OK button in the pop-up.

**Return**

- **Boolean**:

  - True / 1: The pop-up was confirmed with the OK.

  - False / 0: The pop-up was not confirmed with Cancel. This is only a possible option if the Cancel button is visible.

**Example**

This example creates a confirmation pop-up with printed feedback in the Command Line History:

```lua
return function()
    --Creates a pop-up asking to be confirmed and prints a useful text.
    if Confirm("Confirm me", "Tap OK or Cancel", nil, true) then
        Printf("Pop-up result: OK")
    else
        Printf("Pop-up result: Cancel")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_confirm.html>


## MessageBox

**Signature**

```
MessageBox({title:string,[, string:backColor][, integer:timeout (ms)][, boolean:timeoutResultCancel][, integer:timeoutResultID][, string:icon][, string:titleTextColor][, string:messageTextColor] [, boolean:autoCloseOnInput] string:message[, integer:message_align_h(Enums.AlignmentH)][, integer:message_align_v(Enums.AlignmentV)][, integer|lightuserdata:display], commands:{array of {integer:value, string:name[, integer:order]}}, inputs:{array of {string:name, string:value, string:blackFilter, string:whiteFilter, string:vkPlugin, integer:maxTextLength[, integer:order]}}, states:{array of {string:name, boolean:state[, integer:order]}, selectors:{array of {name:string, integer:selectedValue, values:table[, type:integer 0-swipe, 1-radio][, integer:order]} }): {boolean:success, integer:result, inputs:{array of [string:name] = string:value}, states:{array of [string:name] = boolean:state}, selectors:{array of [string:name] = integer:selected-value}}
```

**Help page title:** `MessageBox(table)`

**Description**

The MessageBox Lua function is used to create pop-up message boxes. These can be simple or complex information pop-ups with many different options and user inputs.

The message box contains multiple elements that must be defined in a table. This table is the single input argument to the function.

The elements in the message box are displayed in alphabetical order.

This function is part of the User Interface functions.

**Arguments**

- **Table**:

 The input to the function must be formatted as a table using key-value pairs. The needed elements have default values that will be used if not overwritten. The values can be defined in any order if the key is defined.

 The table can have the following elements:

  - **title**:

 This is the title of the pop-up message box.

  - **titleTextColor**:

 This is the text color for the title text. The value is a number or string that refers to the **UI Colors** in the color theme. See the link below.

  - **backColor**:

 This is the color of the frame or border of the pop-up. The value is a number or string that refers to the **UI Colors** in the color theme. See the link below.

  - **icon**:

 This is an icon that can be shown in the upper left corner of the pop-up. The value can be an integer or a string that refers to the number or name of a texture image (without the file format).

 The icons can be listed by navigating to the texture folder ChangeDestination GraphicsRoot/TextureCollect/Textures and then doing a List command.

 This shows a long list with numbers and names of all the textures in the Command Line History window.

  - **message**:

 This message text string is displayed in the main part of the pop-up. A new line can be created by adding a "\n" in the text.

  - **messageTextColor**:

 This is the text color for the message text. The value is a number or string that refers to the **UI Colors** in the color theme. See the link below.

  - **autoCloseOnInput**:
This option defines if a Please (or Enter) from an input field closes the message box pop-up. The default value is **true**. Setting this to **false** keeps the message box open until it is explicitly closed. See example 5 below.

  - **timeout**:

 The timeout value is an integer that indicates how long the message box is displayed in milliseconds. It will show a countdown timer at the top of the message area. When the countdown runs out, it will close the pop-up.

 When the timeout closes the pop-up, it returns a **success** element with a **true** value (see more about the return below). It was closed "normally".

  - **timeoutResultCancel**:

 This element can change the return to mimic a cancel of the pop-up, which returns a **false** instead of a **true** for the **success **element.

  - **timeoutResultID**:

 When the timeout closes the pop-up with the **success **value = **true**, a special return result can be defined using this, for instance, setting the value to 99 - then 99 is returned when the message box is closed by the timeout.

  - **commands**:

 The commands are buttons at the bottom of the message box pop-up. The input here is a table of objects using the following structure:

    - **value**:

 This integer value will be returned as the result value - see more about the return below.

    - **name**:

 This is a string which will be shown on the button.

  - **inputs**:

 The inputs are user input fields where text or numbers can be input. The input fields will be displayed in alphabetical order.

 The input fields are defined using an table with the following structure (see example 5 below):

    - **name**:

 This is a string value - the text will be shown as a label for the input field.

    - **value**:

 This is a string value - it is a default input value for the input field.

    - **blackfilter**:

 This is a string value - it defines input characters that are not allowed.

    - **whitefilter**:

 This is a string value - it defines which input characters are allowed.

    - **vkPlugin**:

 This is a string value - it is the name of the input pop-up, which is opened if the on-screen keyboard icon () is tapped in the pop-up. Example 5 below has a list of possible vkPlugin values.

    - **maxTextLength**:

 This is an integer value - it defines the maximum number of characters for the input.

  - **states**:

 The states are buttons in the pop-up. State buttons have a small checkbox and can have a true or false state. The buttons will be displayed in alphabetical order.

 The buttons are defined using a table with the following structure (see example 4 below):

    - **name**:

 This is a string value - the text will be shown on the button.

    - **state**:

 This is the initial state of the button checkbox.

  - **selectors**:

 Selector buttons are two different types of buttons. Each type can have a selected value based on a list of available values. The two types are **Swipe **button (type 0) and **Radio **button (type 1).

 The buttons are defined in a table with the following structure:

    - **name**:

 This is a string value - the text will be shown on the button (swipe button) or as a label above the buttons (radio buttons).

    - **selectedValue**:

 This is an integer value - it defines the default selected value

    - **type**:

 This is an integer value - it defines the type of selector button. The options are:

      - **0**:

 This defines the button as a swipe button.

      - **1**:

 This defines the button as a radio button.

    - **values**:

 This is another table containing the different values available for the selector button. Each value element in the table has the following structure: ["string"]=integer

 The string is the name displayed for the value. The integer is the value returned and the one used for **selectedValue** (see above). See example 6 below for an example of use.

The table can contain some or all of the elements described above.

The colors mentioned above can be a string or number value. It refers to a defined UI Color in the color theme, for instance, **"Global.Text"** or **1.27**. See more in the Color Theme topic.

A message box pop-up should have at least a title, message, and *either *a timeout *or* some basic command buttons. See the first three examples below.

**Return**

- **Table**:

 The return from a message box pop-up is formatted in a table. The returned table adjusts to match the elements of the message box. For instance, if there are selection buttons in the message box, then there is another table inside the result table containing the return from the selection buttons. See the examples for details on how to retrieve the results. The table can contain the following elements:

  - **success**:

 This is a boolean - it returns true if the message box was closed by tapping a command button or by a timeout (see **timeoutResultCancel** above for exceptions).

  - **result**:

 This is an integer - it returns the value of the tapped command button or the timeout result ID.

  - **inputs** (only if the message box has inputs fields):

 This is a table with a list of the input fields' name and string value in a key-value pair table.

  - **states** (only if the message box has state buttons):

 This is a table with a list of the state buttons' name and boolean value in a key-value pair table.

  - **selectors** (only if the message box has selector buttons):

 This is a table with a list of the selector buttons' name and integer value in a key-value pair table.

An error is thrown if the message box does not have inputs, states, and selectors, but the script tries to use the table elements.

See the examples to see how to extract the results.

**Example**

There are six different examples demonstrating different elements of the message box. The elements can be combined, but the examples highlight different functions.

Example 1

A simple message box pop-up that shows a single confirm button:

```lua
return function ()
    -- This creates a small pop-up with a single button.
    local returnTable = MessageBox(
        {
            title = "Please confirm This",
            commands = {{value = 1, name = "Confirm"}}
        }
    )

    -- Print the content of the returned table. 
    Printf("Success = "..tostring(returnTable.success))
    Printf("Result = "..returnTable.result)
end
```

Example 2

This example opens a pop-up with some text and two command buttons:

```lua
return function ()
    -- A table with two default buttons for the pop-up
    local defaultCommandButtons = {
        {value = 2, name = "OK"},
        {value = 1, name = "Cancel"}
    }

    -- A table with the elements needed for the pop-up
    local messageTable = {
        icon = "object_smart",
        backColor = "Window.Plugins",
        title = "This is the title",
        message = "This is a message\nThat can have multiple lines",
        commands = defaultCommandButtons,
    }

    -- The creation on the actual pop-up with the result stored in a variable
    local returnTable = MessageBox(messageTable)

    -- Print the content of the returned table
    Printf("Success = "..tostring(returnTable.success))
    Printf("Result = "..returnTable.result)
end
```

Example 3

This example displays a message box for 3 seconds and then closes itself:

```lua
return function ()
    -- This variable contains the table used as argument for the messagebox
    local messageTable = {
        title = "Do not worry",
        message = "This message will self destruct\nGoodbye!",
        timeout = 3000,
        timeoutResultCancel = false,
        timeoutResultID = 99,
    }

    -- This creates the messagebox pop-up and store the return table in a variable
    local returnTable = MessageBox(messageTable)

    -- Print the content of the returned table
    Printf("Success = "..tostring(returnTable.success))
    Printf("Result = "..returnTable.result)
end
```

Example 4

This example adds state buttons to the message box. The buttons are added to a table for a better overview.

```lua
return function ()
    -- A table with two default buttons for the pop-up
    local defaultCommandButtons = {
        {value = 2, name = "OK"},
        {value = 1, name = "Cancel"}
    }
    -- A table with three state buttons 
    -- The buttons will be displayed alphabetically in the pop-up
    local stateButtons = {
        {name = "State B", state = false},
        {name = "State A", state = false},
        {name = "New State", state = false}
    }

    -- A table with the elements needed for the pop-up
    local messageTable = {
        icon = "object_smart",
        backColor = "Window.Plugins",
        title = "This is state buttons",
        message = 'Toggle the states and click "Ok"',
        commands = defaultCommandButtons,
        states = stateButtons,
    }

    -- The creation on the actual pop-up with the result stored in a variable
    local returnTable = MessageBox(messageTable)

    -- Print the content of the returned table
    Printf("Success = "..tostring(returnTable.success))
    Printf("Result = "..returnTable.result)
    
    -- Print a list with the state of the stateButtons
    for name,state in pairs(returnTable.states) do
        Printf("State '%s' = '%s'",name,tostring(state))
    end
end
```

Example 5

This example shows the input fields.

```lua
return function ()
    -- A table with two default buttons for the pop-up
    local defaultCommandButtons = {
        {value = 2, name = "OK"},
        {value = 1, name = "Cancel"}
    }
    -- A table with three input fields
    -- The fields will be displayed alphabetically in the pop-up based on name
    local inputFields = {
        {name = "Numbers Only", value = "1234", whiteFilter = "0123456789", vkPlugin = "NumericInput"},
        {name = "Text Only", value = "abcdef", blackFilter = "0123456789"},
        {name = "Maximum 10 characters", value = "", maxTextLength = 10}
    }
    -- Possible vkPlugin values:
    -- - "TextInput" : same as default - standrd on-screne keyboard
    -- - "TextInputNumOnly" : text input but only with number buttons
    -- - "TextInputNumOnlyRange" : text input but only with number and related range buttons 
    -- - "TextInputTimeOnly" : text input styled for time input - includes buttons for time values
    -- - "NumericInput" : general number input
    -- - "CueNumberInput" : number input styled for cue number
    -- - "RelCueNumberInput" : number input with the relative "delta" button
    -- - "IP4Prefix" : designed for inputting an IPv4 address allowing CIDR notation

    -- A table with the elements needed for the pop-up
    local messageTable = {
        icon = "object_smart",
        backColor = "Window.Plugins",
        title = "This is input fields",
        message = 'Change the values in the input fields and click "Ok"',
        commands = defaultCommandButtons,
        inputs = inputFields,
        autoCloseOnInput = false
    }

    -- The creation on the actual pop-up with the result stored in a variable
    local returnTable = MessageBox(messageTable)

    -- Print the content of the returned table
    Printf("Success = "..tostring(returnTable.success))
    Printf("Result = "..returnTable.result)
    -- Print a list with the values of the input fields
    for name,value in pairs(returnTable.inputs) do
        Printf("Input '%s' = '%s'",name,tostring(value))
    end
end
```

Example 6

This example shows the different selector buttons.

```lua
return function ()
    -- A table with two default buttons for the pop-up
    local defaultCommandButtons = {
        {value = 2, name = "OK"},
        {value = 1, name = "Cancel"}
    }
    -- A table with selector buttons
    -- The buttons will be displayed alphabetically in the pop-up based on name
    local selectorButtons = {
        { name="Swipe Selector", selectedValue=1, type=0, values={["Swipe1"]=1,["Swipe2"]=2}},
        { name="Radio Selector", selectedValue=2, type=1, values={["Radio1"]=1,["Radio2"]=2}},
        { name="Another Radio", selectedValue=3, type=1, values={["Radio3"]=3,["Radio4"]=4}}
    }

    -- State button to show grouping with swipe Selector button
    local stateButton = {
        {name = "State Button", state = false},
    }
    -- A table with the elements needed for the pop-up
    local messageTable = {
        icon = "object_smart",
        backColor = "Window.Plugins",
        title = "This is input fields",
        message = 'Change the values in the input fields and click "Ok"',
        commands = defaultCommandButtons,
        states = stateButton,
        selectors = selectorButtons,
    }

    -- The creation on the actual pop-up with the result stored in a variable
    local returnTable = MessageBox(messageTable)

    -- Print the content of the returned table
    Printf("Success = "..tostring(returnTable.success))
    Printf("Result = "..returnTable.result)
    -- Print a list with the values of the selection buttons
    for name,value in pairs(returnTable.selectors) do
        Printf("Input '%s' = '%s'",name,tostring(value))
    end
    -- Print a list with the state of the stateButton
    for name,state in pairs(returnTable.states) do
        Printf("State '%s' = '%s'",name,tostring(state))
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_messagebox.html>


## GetDisplayByIndex

**Signature**

```
GetDisplayByIndex(integer:display_index): light_userdata:display_handle
```

**Help page title:** `GetDisplayByIndex(integer)`

**Description**

The **GetDisplayByIndex **Lua function returns a handle to the display object matching the provided index number.

**Arguments**

- **Integer**:

            This function needs an index number for one of the displays.

**Return**

- **Handle**:

            The returned handle to the display object.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Get the index number for "Display 1"
    local displayIndex = GetDisplayCollect()["Display 1"].INDEX
    -- return error text in case the index number is nil
    if displayIndex == nil then
        ErrPrintf('Something went wrong. It appears that there is no "display 1"')
        return
    end
    -- Dump all information about the display with the index number
    Printf("=============== START OF DUMP ===============")
    GetDisplayByIndex(displayIndex):Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getdisplaybyindex.html>


## GetRemoteVideoInfo

**Signature**

```
GetRemoteVideoInfo(nothing): integer:wingID, boolean:isExtension
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetUIObjectAtPosition

**Signature**

```
GetUIObjectAtPosition(integer:display_index, {x=integer:x_position,y=integer:y_position}): light_userdata:handle to UI object or nil
```

**Help page title:** `GetUIObjectAtPosition(integer, table)`

**Description**

The **GetUIObjectAtPosition** Lua function returns the handle of the UI Object at a specified position on a specified display.

**Arguments**

- **Integer**:
The integer should be the index number of the display with the UI object.
- **Table**:
  The table must have two elements with the following keys:

  - x: This is the X position on the display. The value must be a number indicating the desired pixel position. It is counted from the left side of the display.

  - y: This is the Y position on the display. The value must be a number indicating the desired pixel position. It is counted from the top of the display.

**Return**

- **Handle** | **nil**:
  If a UI object is at the provided position, then the handle to the object is returned. Otherwise, it returns nil.

**Example**

This example prints the Dump of the UIObject at a specific position on display 1. It also uses the DrawPointer function to draw a red pointer at the position.

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

The DrawPointer function draws a red pointer at a display. Learn more about it in the DrawPointer() topic.

```lua
return function()
    -- Get the index number for "Display 1"
    local displayIndex = GetDisplayCollect()["Display 1"].INDEX
    -- Create a table with X and Y position
    local positionTable = {}
    positionTable.x = 1000
    positionTable.y = 500
    -- Get the UI object handle
    local uiObjectAtPositionHandle = GetUIObjectAtPosition(displayIndex,positionTable)
    -- Dump all information about the display with the index number if not nil
    if uiObjectAtPositionHandle == nil then
        Printf("The returned value was not a valid handle.")
        return
    end
    -- Draw a pointer at the posiiton for 5 seconds
    DrawPointer(displayIndex,positionTable,5000)
    --Dump of the UIObject
    Printf("=============== START OF DUMP ===============")
    uiObjectAtPositionHandle:Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getuiobjectatposition.html>


## DrawPointer

**Signature**

```
DrawPointer(integer:display_index, {x=integer:x_position,y=integer:y_position}, integer:duration in ms)): nothing
```

**Help page title:** `DrawPointer(integer,table[,integer])`

**Description**

The **DrawPointer **function draws a red pointer on the display. There can only be one pointer at a time on each station.

**Arguments**

- **Integer**:

 This integer is the display index where the pointer should be drawn.

- **Table**:

 This key-value table must have 'x' and 'y' keys with values indicating a position on the display. See the example below.

- **Integer** (optional):

 This optional integer defines a duration for the pointer in milliseconds. It fades out. If a duration is not set, then it stays visible.

**Return**

This function does not return anything.

**Example**

This example draws a pointer on display 1 for 5 seconds:

```lua
return function()
    --Set a display index
    local displayIndex = 1
    --Create and set the position in a table
    local position = {}
    position.x = 150
    position.y = 25
    --Set a 5 seconds duration - in milliseconds
    local duration = 5000
    --Draw the actual pointer
    DrawPointer(displayIndex,position,duration)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_drawpointer.html>


## WaitObjectDelete

**Signature**

```
WaitObjectDelete(light_userdata:handle to UIObject[, number:seconds to wait]): boolean:true on success, nil on timeout
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetFocus

**Signature**

```
GetFocus(nothing): light_userdata:display_handle
```

**Help page title:** `GetFocus()`

**Description**

The **GetFocus** Lua function returns a handle to the object that currently has focus in the UI.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The returned handle to the object.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- This example dumps all information about the object who currently got focus.
    Printf("=============== START OF DUMP ===============")
    GetFocus():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getfocus.html>


## GetFocusDisplay

**Signature**

```
GetFocusDisplay(nothing): light_userdata:display_handle
```

**Help page title:** `GetFocusDisplay()`

**Description**

The **GetFocusDisplay** Lua function returns a handle to the display object that currently has focus in the UI.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the display object.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- This example dumps all information about the display object who currently got focus.
    Printf("=============== START OF DUMP ===============")
    GetFocusDisplay():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getfocusdisplay.html>


## GetDisplayCollect

**Signature**

```
GetDisplayCollect(nothing): light_userdata:handle to DisplayCollect
```

**Help page title:** `GetDisplayCollect()`

**Description**

The **GetDisplayCollect** Lua function returns a handle to the DisplayCollect object.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The returned handle to the DisplayCollect.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- This example dumps all information about the DisplayCollect object.
    Printf("=============== START OF DUMP ===============")
    GetDisplayCollect():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getdisplaycollect.html>


## FindBestFocus

**Signature**

```
FindBestFocus([light_userdata:handle]): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## FindNextFocus

**Signature**

```
FindNextFocus([bool:backwards(false)[, int(Focus::Reason):reason(UserTabKey)]]): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## CloseAllOverlays

**Signature**

```
CloseAllOverlays(nothing): nothing
```

**Help page title:** `CloseAllOverlays()`

**Description**

The **CloseAllOverlays** function closes any pop-ups or menus (overlays) open on any screen.

**Arguments**

This function does not accept any arguments.

**Return**

This function does not return anything.

**Example**

This example simply closes any overlay.

```lua
return function()
    CloseAllOverlays()
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_closealloverlays.html>


## GetTopModal

**Signature**

```
GetTopModal(nothing): light_userdata:handle to top modal overlay
```

**Help page title:** `GetTopModal()`

**Description**

The **GetTopModal** Lua function returns a handle for the modal at the top. Modal is the internal name for pop-ups that interrupt the system's normal operation. A modal blocks other UI elements from being used while it is open.

For example, when opening a window's settings pop-up, it is not possible to use the command line. The settings pop-up is a modal. Modals can also be identified by the rest of the UI, which darkens a bit when it is open.

**Arguments**

This function does not have any arguments.

**Return**

- **Handle** | **nil**:

 The function returns a handle to the top modal UI object if there is one.

**Example**

This example uses the Dump() function to show information about the StagePopup selection pop-up.

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Open a Modal / Pop-up.
    Cmd('Menu "StagePopup"')
    -- Add a small wait.
    coroutine.yield(0.5)
    -- Get the handle for the modal / pop-up. 
    local modalHandle = GetTopModal()
    -- If there is a handle then dump all information else print en error feedback.
    if modalHandle ~= nil then
        Printf("=============== START OF DUMP ===============")
        modalHandle:Dump()
        Printf("================ END OF DUMP ================")
    else
        ErrPrintf("The Modal UI object could not be found.")
    end
    -- Close the modal / pop-up by pressing the Escape key.
    Keyboard(1,'press','Escape')
    Keyboard(1,'release','Escape')
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_gettopmodal.html>


## GetTopOverlay

**Signature**

```
GetTopOverlay(integer:display_index): light_userdata:handle to top overlay on the display
```

**Help page title:** `GetTopOverlay()`

**Description**

The **GetTopOverlay** Lua function returns a handle for the overlay at the top of the display with the provided index number. Overlay is the internal name for what is called pop-ups or menus in the rest of this manual.

**Arguments**

This function does not have any arguments.

**Return**

- **Handle** | **nil**:

 The function returns a handle to the top overlay UI object if there is one.

**Example**

This example uses the Dump() function to show information about the MenuSelector pop-up - it is the one opening when pressing the Menu key.

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Open the MenuSelector overlay.
    Cmd('Menu "MenuSelector')
    -- Add a small delay.
    coroutine.yield(0.5)
    -- Get the handle for the overlay on the display with index 1.
    local overlayHandle = GetTopOverlay(1)
    -- Add a small delay.
    coroutine.yield(0.5)
    -- Close the MenuSelector overlay.
    Cmd('Menu "MenuSelector')
    -- Check if there is a handle and print appropriate feedback.
    if overlayHandle ~= nil then
        Printf("=============== START OF DUMP ===============")
        overlayHandle:Dump()
        Printf("================ END OF DUMP ================")
    else
        ErrPrintf("The Overlay UI object could not be found.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_gettopoverlay.html>


## WaitModal

**Signature**

```
WaitModal([number:seconds to wait]): handle to modal overlay or nil on failure(timeout)
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SetBlockInput

**Signature**

```
SetBlockInput(boolean:block[, boolean:show_info]): nothing
```

**Help page title:** `SetBlockInput(boolean)`

**Description**

The **SetBlockInput **function is an internal function used during the system tests. It stops input from USB-connected keyboards and mouse. The built-in keyboard on some models is internally connected using a USB connection, which is also blocked by this function. The block affects the station where the function is executed.

            **Important:**

Should the station be blocked and cannot be unblocked using a new Lua command, then pressing the keyboard keys A, S, D, and F simultaneously unblock the input again.

**Arguments**

- **Boolean**:

 The boolean indicates if the stations' input should be blocked or unblocked.

  - **true** (or 1): The station input is blocked.

  - **false** (or 0): The station input is unblocked.

**Return**

This function does not return anything.

**Example**

This example blocks mouse and keyboard input for 10 seconds:

```lua
return function()
    -- Set a variable for yield time in seconds
    yieldTime = 10
    -- Set the block to true
    SetBlockInput(true)
    -- Wait the [yieldtime]
    coroutine.yield(yieldTime)
    -- Unblock the station
    SetBlockInput(false)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_setblockinput.html>


## GetBlockInput

**Signature**

```
GetBlockInput(nothing): boolean:block_input
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## FindTexture

**Signature**

```
FindTexture(string:texture name): light_userdata:handle to texture found
```

**Help page title:** `FindTexture(string)`

**Description**

The **FindTixture** Lua function returns a handle to the texture matching the input text string - if the texture exists.

**Arguments**

- **String**:

            The text string must be the name of the texture without the file type. See the example below.

**Return**

- **Handle** | **nil**:

            The function returns the texture handle or nil if it does not exist.

**Example**

This example prints the information about the "button" texture. The example uses the Dump() function.

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- Set a texture name.
    local textureName = "button"
    -- Get the handle of the texture.
    local textureHandle = FindTexture(textureName)
    -- Check if textureHandle returned something and provide feedback.
    if textureHandle == nil then
        ErrPrintf("Texture does not exist.")
    else
        Printf("=============== START OF DUMP ===============")
        FindTexture(textureName):Dump()
        Printf("================ END OF DUMP ================")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_findtexture.html>


## GetScreenContent

**Signature**

```
GetScreenContent(light_userdata:handle to ScreenConfig): light_userdata:handle
```

**Help page title:** `GetScreenContent(handle)`

**Description**

The **GetScreenContent** Lua function returns a handle to the screen content based on a provided handle to a screen configuration.

**Arguments**

- **Handle**:
This must be a handle to a screen configuration.

**Return**

- **Handle**:

 The returned handle to the screen content.

**Example**

This example prints the data connected to the screen content handle. It uses the CurrentScreenConfig() and Dump() functions:

The **CurrentScreenConfig** Lua function returns a handle to the current users' screen configuration. Learn more in the CurrentScreenConfig topic.

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Create a handle for the current screen configuration.
    local myCurrentScreenConfig = CurrentScreenConfig()
    -- Create a handle for the screen content based on the screen configuration.
    local myScreenContent = GetScreenContent(myCurrentScreenConfig)
    -- Print the Dump of the handle.
    Printf("=============== START OF DUMP ===============")
    myScreenContent:Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getscreencontent.html>


## FSExtendedModeHasDots

**Signature**

```
FSExtendedModeHasDots(light_userdata:handle to UIGrid (or derived), {r, c}:cell): boolean
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)

---

[Back to index](README.md)
