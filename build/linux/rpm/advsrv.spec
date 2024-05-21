Name:           advsrv
Version:        ${version}
Release:        ${release}%{?dist}
Summary:        An advertisements managment package.
ExclusiveArch:  x86_64
License:        MIT

%description
An advertisements managment package.
Build number ${buildNumber}.

Requires(pre): /usr/sbin/useradd, /usr/bin/getent

%pre
/usr/bin/getent passwd advsrv 2>/dev/null || /usr/sbin/useradd -r -d /path/to/program -s /sbin/nologin advsrv

%post
if [ ! -f /etc/advsrv/config.yaml ];
then 
  mkdir %{_sysconfdir}/advsrv
  %{_sbindir}/advsrv generate > %{_sysconfdir}/advsrv/config.yaml
  chown advsrv:advsrv %{_sysconfdir}/advsrv/config.yaml
fi

%preun
if $(systemctl -q is-active advsrv.service);
then 
  systemctl stop advsrv.service
fi

%postun
rm -rf %{_sysconfdir}/advsrv

%install

mkdir -p %{buildroot}/%{_sbindir}
install -m 0755 %{name} %{buildroot}/%{_sbindir}

mkdir -p %{buildroot}/%{_sysconfdir}/advsrv/

%files
%{_sbindir}/advsrv