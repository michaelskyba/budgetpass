# budgetpass
budgetpass: a simple albeit likely insecure password manager written in Go

## Overview
budgetpass stores each of your ("local") passwords in encrypted files, where
the name of the file is the name you gave that password. The encryption uses
AES, where your master password is the key. So, unless somebody has your master
password, they won't be able to decrypt your password files and hence will not
have access to any of your passwords.

The directory in which these password files are stored is
``$HOME/.local/share/bpass`` by default, but can be modified
via the ``$BP_HOME`` variable. ``$BP_HOME`` will be used in the examples.

## Spaces Disclaimer
From my (limited) testing, everything seems to break when spaces are involved,
so avoid them.

## Commands
1. ``echo "<master password>\n<local password>" | bpass new <password name>`` - create a new password.
2. ``echo "<master password>" | bpass get <password name>`` - retrieve a password.

## Example extensions
- List passwords - ``ls $BP_HOME``
- Delete a password - ``rm $BP_HOME/<password name>``
- Rename a password - ``cd $BP_HOME && mv <old name> <new name>``
- Update a password - ``rm $BP_HOME/<password name> && bpass new <password name>``

As you can see, the point is for budgetpass to only provide the
encryption/decryption logic. The rest of what might be considered basic
password manager functionality is left for your scripts, as you want it.

## Errors
If, when using ``bpass new``, you get some sort of ``open ...: no such file or directory``
error, run ``mkdir -p ${BP_HOME:-$HOME/.local/share/bpass}``.

When using ``bpass get``, an error like this:
```
open /home/michael/.local/share/bpass/michael: no such file or directory
panic: runtime error: slice bounds out of range [:12] with capacity 0

goroutine 1 [running]:
main.main()
        /home/michael/s/budgetpass/bpass.go:105 +0x10f5
```
indicates that the password you're trying to access does not exist.

## Installation
Clone the repo, run ``go build bpass.go``, and then copy the resulting
``bpass`` binary into your ``$PATH``.
