%global commit      f5c47f460f687a4e32e9b33ff0c1e67445dcb59b
%global shortcommit %(c=%{commit}; echo ${c:0:7})

Name:           securepass-tools
Version:        0.5
Release:        1%{?dist}
Summary:        SecurePass Tools 
License:        GPLv2+
URL:            https://github.com/garlsecurity/securepassctl
#Source0:        https://github.com/garlsecurity/securepassctl/archive/v%{version}/securepass-tools-v%{version}.tar.gz  
Source0:        https://github.com/garlsecurity/securepassctl/archive/%{commit}/securepassctl-%{shortcommit}.tar.gz


BuildRequires:  gcc
BuildRequires:  golang >= 1.2-7

#BuildRequires:  golang(github.com/gorilla/mux) >= 0-0.13

%description
# include your full description of the application here.

%prep
%setup -q -n securepassctl-%{commit}

# many golang binaries are "vendoring" (bundling) sources, so remove them. Those dependencies need to be packaged independently.
rm -rf vendor

%build
export GOPATH=$(pwd):%{gopath}
make deps
GOOS=linux GOARCH=amd64 go build -v -ldflags="-X=main.Version=%{version}" -o build/spctl .


%install
install -d %{buildroot}%{_bindir}
install -p -m 0755 ./build/linux/%{_arch}/spctl %{buildroot}%{_bindir}/spctl

%files
%defattr(-,root,root,-)
#%doc AUTHORS CHANGELOG.md CONTRIBUTING.md FIXME LICENSE MAINTAINERS NOTICE README.md
%{_bindir}/spctl

%changelog
* Tue Jul 01 2014 Jill User <jill.user@fedoraproject.org> - 1.0.0-6
- package the example-app
