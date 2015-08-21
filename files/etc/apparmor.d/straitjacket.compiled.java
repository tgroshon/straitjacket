#include <tunables/global>

profile straitjacket/compiled/java {
  #include <abstractions/base>
  #include <abstractions/straitjacket-base>
  #include <abstractions/straitjacket-jvm>

  /etc/java-8-openjdk/jvm-amd64.cfg r,
}
