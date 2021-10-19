Name:           serial
Version:        0.0.0
Release:        1%{?dist}
Summary:        Serial is a modern command-line serial port tool designed with both humans and machines in mind.

License:        GPLv3
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  golang

Provides:       %{name} = %{version}

%description
Serial is a modern command-line serial port tool designed with both humans and machines in mind.

%global debug_package %{nil}

%prep
%autosetup

%build
bash contrib/build/build.sh

%install
install -Dpm 0755 build/%{name} %{buildroot}%{_bindir}/%{name}

%check
go test ./...

#%post
#%systemd_post %{name}.service

#%preun
#%systemd_preun %{name}.service

%files
%{_bindir}/%{name}

%changelog
* Wed May 19 2021 John Doe - 1.0-1
- First release%changelog