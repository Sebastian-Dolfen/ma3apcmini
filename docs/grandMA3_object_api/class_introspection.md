# Class Introspection

Determining what kind of object you have and what kind of children it contains.

## Contents

- [GetClass](#getclass)
- [GetChildClass](#getchildclass)

---

## GetClass

**Signature:**

```
GetClass(light_userdata:handle): string:class_name
```

**Summary:** The GetClass function returns a string with information about the class for the object.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.

**Returns:**

- String : The function returns the text string with the name of the object's class.

**Used by APC plugin for:** guarding code paths so we only call Sequence-specific methods on Sequence handles.

**Example:**

This example prints the class name of the selected sequence.

```lua
return function()
    -- Gets the class name of the selected sequence.
    Printf("The class name is " .. SelectedSequence():GetClass())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_getclass.html>

---

## GetChildClass

**Signature:**

```
GetChildClass(light_userdata:handle): string:class_name
```

**Summary:** The GetChildClass function returns a string with the name of the class of the object's children.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object

**Returns:**

- String : The function returns a text string with the name of the class of the object's children.

**Used by APC plugin for:** discovering what kind of children to expect before iterating (Cue vs Executor).

**Example:**

This example prints the class name of the selected sequences' children.

```lua
return function()
    -- Gets the class name of children of the selected sequence.
    Printf("The class name is " .. SelectedSequence():GetChildClass())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_getchildclass.html>

---

[Back to index](README.md)
