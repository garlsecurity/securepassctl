%global commit      d227bb8feba8b0030e463dc9ee4ccac78655a1ed
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

# pull in golang libraries by explicit import path, inside the meta golang()
BuildRequires:  golang(github.com/gorilla/mux) >= 0-0.13

%description
# include your full description of the application here.

%prep
%setup -q -n securepassctl-%{version}

# many golang binaries are "vendoring" (bundling) sources, so remove them. Those dependencies need to be packaged independently.
rm -rf vendor

%build
# set up temporary build gopath, and put our directory there
mkdir -p ./_build/src/github.com/example
ln -s $(pwd) ./_build/src/github.com/example/app

export GOPATH=$(pwd)/_build:%{gopath}
go build -o example-app .

%install
install -d %{buildroot}%{_bindir}
install -p -m 0755 ./example-app %{buildroot}%{_bindir}/example-app

%files
%defattr(-,root,root,-)
#%doc AUTHORS CHANGELOG.md CONTRIBUTING.md FIXME LICENSE MAINTAINERS NOTICE README.md
%{_bindir}/example-app

%changelog
* Tue Jul 01 2014 Jill User <jill.user@fedoraproject.org> - 1.0.0-6
- package the example-app
