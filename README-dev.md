# Git Tracking

The main repo is github.com/s3packer/s3p.git, but it is also tracked in a separate remote repo on a Tailscale network.

`git clone git@github.com:orme292/s3p.git`
`git remote set-url origin --push --add http://onedev.siberian-barley.ts.net/s3p-backup`
`git remote set-url origin --push --add git@github.com:orme292/s3p.git`
