# Input simulation (keyboard / mouse / touch)

Synthesize keyboard, mouse and touch events at either screen or object scope, and read system time.

## Functions

- [`Keyboard`](#keyboard)
- [`Mouse`](#mouse)
- [`Touch`](#touch)
- [`MouseObj`](#mouseobj)
- [`TouchObj`](#touchobj)
- [`KeyboardObj`](#keyboardobj)
- [`Time`](#time)

---


## Keyboard

**Signature**

```
Keyboard(integer:display_index, string:type('press', 'char', 'release')[ ,string:char(for type 'char') or string:keycode, boolean:shift, boolean:ctrl, boolean:alt, boolean:numlock]): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## Mouse

**Signature**

```
Mouse(integer:display_index, string:type('press', 'move', 'release')[ ,string:button('Left', 'Middle', 'Right' for 'press', 'release') or integer:abs_x, integer:abs_y)]): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## Touch

**Signature**

```
Touch(integer:display_index, string:type('press', 'move', 'release'), integer:touch_id, integer:abs_x, integer:abs_y): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## MouseObj

**Signature**

```
MouseObj(nothing): light_userdata:handle
```

**Help page title:** `MouseObj()`

**Description**

The **MouseObj** function returns the handle to the first found mouse object.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The function returns the handle to the mouse object.

**Example**

This example prints the information of the mouse object. Is uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Print all informatin about the MouseObj object
    Printf("=============== START OF DUMP ===============")
    MouseObj():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_mouseobj.html>


## TouchObj

**Signature**

```
TouchObj(nothing): light_userdata:handle
```

**Help page title:** `TouchObj()`

**Description**

The **TouchObj** function returns the handle to the first found touch object.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The function returns the handle to the touch object.

**Example**

This example prints information about the touch object using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Print all informatin about the TouchObj object
    Printf("=============== START OF DUMP ===============")
    TouchObj():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_touchobj.html>


## KeyboardObj

**Signature**

```
KeyboardObj(nothing): light_userdata:handle
```

**Help page title:** `KeyboardObj()`

**Description**

The **KeyboardObj** function returns the handle to the first found keyboard object.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The function returns the handle to the keyboard object.

**Example**

This example prints the information of the keyboard object. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Print all informatin about the KeyboardObj object
    Printf("=============== START OF DUMP ===============")
    KeyboardObj():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_keyboardobj.html>


## Time

**Signature**

```
Time(nothing): integer:time
```

**Help page title:** `Time()`

**Description**

The **Time** function returns the time (in seconds) the station has been on, as a number (float). It is basically a stopwatch that starts when the grandMA3 application starts. It is not the current time of day or the session online time.

**Arguments**

This function does not accept any arguments.

**Return**

- **Number**:

            The returned number (float) is the on-time for the station.

**Example**

This example finds, formats, and prints the time.

```lua
return function()
    -- Get the current time
    local time = Time()

    --Calculate the different elements
    local days = math.floor(time/86400)
    local hours = math.floor((time % 86400)/3600)
    local minutes = math.floor((time % 3600)/60)
    local seconds = math.floor(time % 60)

    --Print the result
    Printf("The time is %d:%02d:%02d:%02d", days, hours, minutes, seconds)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_time.html>

---

[Back to index](README.md)
