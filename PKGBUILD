pkgname=gigfun
pkgver=1.0
pkgrel=2
pkgdesc='Kludge for a Gigabyte videocard'
arch=('x86_64')
url='https://git.vimusov.space/me/gigfun'
license=('GPL')
depends=('nvidia-settings' 'nvidia-utils')
makedepends=('go' 'make')
source=("${pkgname}.go" makefile service)
md5sums=('SKIP' 'SKIP' 'SKIP')

build()
{
    make -C "$srcdir"
}

package()
{
    make -C "$srcdir" DESTDIR="$pkgdir" install
    install -D --mode=0644 "$srcdir"/service "$pkgdir"/usr/share/$pkgname/service.example
}
