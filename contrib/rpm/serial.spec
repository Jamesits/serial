Name:           serial
Version:        0.0.1
Release:        1%{?dist}
Summary:        Serial is a modern command-line serial port tool designed with both humans and machines in mind.

License:        GPLv3
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  golang
BuildRequires:  systemd-rpm-macros

Provides:       %{name} = %{version}

%description
A simple web app

%global debug_package %{nil}

%prep
%autosetup


%build
bash contrib/build/build.sh


%install
install -Dpm 0755 build/%{name} %{buildroot}%{_bindir}/%{name}
#install -Dpm 0755 config.json %{buildroot}%{_sysconfdir}/%{name}/config.json
#install -Dpm 644 %{name}.service %{buildroot}%{_unitdir}/%{name}.service

%check
# go test should be here... :)
true

#%post
#%systemd_post %{name}.service

#%preun
#%systemd_preun %{name}.service

%files
#%dir %{_sysconfdir}/%{name}
%{_bindir}/%{name}
#%{_unitdir}/%{name}.service
#%config(noreplace) %{_sysconfdir}/%{name}/config.json


%changelog
* Wed May 19 2021 John Doe - 1.0-1
- First release%changelog