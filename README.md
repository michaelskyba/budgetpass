# BudgetPass
budgetpass: a simple albeit likely insecure password manager written in Go

## Usage
### Basic Explanation
budgetpass stores each of your passwords in an encrypted file, where the name
of the file is the name you gave that password. The encryption uses AES, where
your master password is the key. So, unless somebody has your master password,
they won't be able to decrypt your password file, and hence will not have access
to any of your passwords.

The directory in which these password files are stored is
``$HOME/.local/share/bpass`` by default, but can be modified
via the ``$BP_HOME`` variable.

### Commands
- ``bpass new <password name>`` - create a new password. You will need quotation
marks if your name contains spaces, but I would recommend against doing this,
because filenames with spaces are ugly. Instead, use dashes, underscores, or camelCase.
- ``bpass get <password name>`` - retrieve a password. Again, spaces require quotes.
If you want to get a password without a prompt, possibly because you're using some kind of frontend,
I think ``pass=$(echo <MASTERPASSWORD> | bpass get <PASSWORDNAME>) ; echo ${pass##* }``
should generally work.

### Errors
If, when using ``bpass new``, you get some sort of ``open ...: no such file or directory``
error, run ``mkdir -p "${BP_HOME:-$HOME/.local/share/bpass}"`` in a terminal.

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

