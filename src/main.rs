
fn main() {
let help = r##"stw  [COMMAND] [OPTION ...] file ...

COMMANDS:
    stw             Stow the package names
    delete           Unstow the package names

    import           Import existing files into stow package from target
    export           Takes files from source path and copis them to the target path

    restow           Restow (like stow -D followed by stow -S)
    help             Show this help
    version          Show plow version number

    try              Donwloads repo from source and temporary replaces current semlinks
                     with new that poitns to the downloaded repo

OPTIONS:
    -f, --from=DIR           Set stow dir to DIR (default is current dir)
    -t, --to=DIR             Set target to DIR (default is parent of stow dir)
    -i  --ignore=REGEX       Ignore files ending in this Perl regex
        --force              Force stowing files beginning with this Perl regex
                             if the file is already stowed to another package
    -d, --dotfiles           Enables special handling for dotfiles that are
                             Stow packages that start with "dot-" and not "."
    --prefix                 Prefix
    --suffix                 Suffix
    -n, --no,                Log and don't write changes to filesystem
    -v, --verbose[=N]        Increase verbosity (levels are from 0 to 5;
                             -v or --verbose adds 1; --verbose=N sets level)
                             "##;
    println!("{}", help);
}
