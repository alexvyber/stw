package main

import (
	"fmt"
)

func main() {
	// TODO:
	// traverse source dir tree
	// traverse destination dir tree
	//
	// keep tack of stows(save files that is being changed)
	//
	// cases:
	// 1. happy path: not file form source dir exists in destination dir
	// 2. if link exists only in source dir create link that points to same resource
	// 3. if link exists in dest and points to the recourse inside source dir, overwrite
	// 4. if link exists in dest and points to the recourse inside source dir, error
	// 5. if file exists in dest, error
	// 5.1. if overwrite flag provided - cp file|link into backup folder and craete link to file in source dir
	// 5.2. if adapt flag|command, cp source file|link into backup folder, craete dest file into source, create link to source dir
	// 6. always find highest tree point of entry: if source dir has one/two/three and des has one/, create symlink to dir two
	// 7. when delete, delete only links that point to source dir
	// 8. when adopt, save backup
	// 9. when restore, restore last saved in backup state
	// 10. if partial provided, allow to complete with errors
	//

	help := `
COMMANDS:

delete                  remove symlinks to source dir files from dest dir
adopt                   import existing files from dest dir into source dir
help                    show this help
version                 show stw version number
restore                 restore state before last invocation

OPTIONS:

-s DIR, --source=DIR    set source dir to DIR (default is current dir)
-t DIR, --target=DIR    set target dir to DIR (default is parent of stow dir)

-p, --partial           allow ignore errors with separate files and continue
-f, --files             create symlinks only for the provided files
-i, --ignore            ignore files
-d, --defer             ff symlinks exist, do nothing
-v, --verbose           verbose

--overwrite             remove file in dest dir and create symlink to source dir

--dry                   do not make any filesystem changes
`

	fmt.Println(help)

}
