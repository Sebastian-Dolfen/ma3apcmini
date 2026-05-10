# Messaging (Lua queues)

Open named Lua message queues and post messages between Lua tasks or stations.

## Functions

- [`OpenMessageQueue`](#openmessagequeue)
- [`CloseMessageQueue`](#closemessagequeue)
- [`SendLuaMessage`](#sendluamessage) [APC]

---


## OpenMessageQueue

**Signature**

```
OpenMessageQueue(string:queue name): boolean:success
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## CloseMessageQueue

**Signature**

```
CloseMessageQueue(string:queue name): boolean:success
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SendLuaMessage

**Signature**

```
SendLuaMessage(string:ip/station, string:channel name, table:data): boolean:success
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

**Used by APC plugin for:** fires a Lua message to a peer station - used if the bridge runs on another MA3 node.

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)

---

[Back to index](README.md)
