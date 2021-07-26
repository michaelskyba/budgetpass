# BudgetPass
budgetpass: a simple albeit likely insecure password manager written in Go

## Overview
budgetpass stores each of your passwords in an encrypted file, where the name
of the file is the name you gave that password. The encryption uses AES, where
your master password is the key. So, unless somebody has your master password,
they won't be able to decrypt your password file, and hence will not have access
to any of your passwords.

The directory in which these password files are stored is
``$HOME/.local/share/bpass`` by default, but can be modified
via the ``$BP_HOME`` variable.

## Spaces Disclaimer
From my (limited) testing, everything seems to break when spaces are involved,
so attempt to avoid them, if possible.

## Commands
- ``bpass new <password name>`` - create a new password.
- ``bpass get <password name>`` - retrieve a password. If you want to get a
password without a prompt, possibly because you're using some kind of frontend,
I think ``pass=$(echo <master password> | bpass get <password name>) ; echo ${pass##* }``
should generally work.

## Example Extensions
- List passwords - ``ls ${BP_HOME:-$HOME/.local/share/bpass}``
- Delete a password - ``rm ${BP_HOME:-$HOME/.local/share/bpass}/<password name>``
- Rename a password - ``cd ${BP_HOME:-$HOME/.local/share/bpass} && mv <old name> <new name> && cd -``
- Update a password - ``rm ${BP_HOME:-$HOME/.local/share/bpass}/<password name> && bpass new <password name>``

## Errors
If, when using ``bpass new``, you get some sort of ``open ...: no such file or directory``
error, run ``mkdir -p ${BP_HOME:-$HOME/.local/share/bpass}`` in a terminal.

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

