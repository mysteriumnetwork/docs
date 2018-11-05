## [0.5-rc](https://github.com/mysteriumnetwork/node/tree/0.5-rc) (2018-10-24)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.4.0...0.5-rc)

**Implemented enhancements:**

- Stale PR should be marked in a 2 weeks [\#436](https://github.com/mysteriumnetwork/node/issues/436)
- Semantic version for dev releases [\#469](https://github.com/mysteriumnetwork/node/pull/469) ([soffokl](https://github.com/soffokl))

**Fixed bugs:**

- Pre-builded images does not provide correct version information.  [\#441](https://github.com/mysteriumnetwork/node/issues/441)
- Example docker run command in README fails [\#439](https://github.com/mysteriumnetwork/node/issues/439)
- Tequilapi http client never times-out [\#386](https://github.com/mysteriumnetwork/node/issues/386)
- Updated chzyer/readline to fix CLI quit command. \#462 [\#477](https://github.com/mysteriumnetwork/node/pull/477) ([soffokl](https://github.com/soffokl))
- Added token usage for Github API calls [\#470](https://github.com/mysteriumnetwork/node/pull/470) ([soffokl](https://github.com/soffokl))
- Fixed build in version in docker images. \#441 [\#468](https://github.com/mysteriumnetwork/node/pull/468) ([soffokl](https://github.com/soffokl))

**Closed issues:**

- Quiting the cli hangs instead of quitting [\#462](https://github.com/mysteriumnetwork/node/issues/462)
- Promise check should be available through --experiment-promise-check flag only [\#456](https://github.com/mysteriumnetwork/node/issues/456)
- Semantic version for dev releases [\#434](https://github.com/mysteriumnetwork/node/issues/434)
- Endpoint `/proposals?fetchQuality=true` merges Discovery + Quality Oracle data [\#429](https://github.com/mysteriumnetwork/node/issues/429)
- Make 'myst daemon' as separate command [\#402](https://github.com/mysteriumnetwork/node/issues/402)
- Paid identity POC [\#334](https://github.com/mysteriumnetwork/node/issues/334)

**Merged pull requests:**

- myst cli does not start new daemon [\#475](https://github.com/mysteriumnetwork/node/pull/475) ([soffokl](https://github.com/soffokl))
- Enable easier node running in development environment [\#474](https://github.com/mysteriumnetwork/node/pull/474) ([Waldz](https://github.com/Waldz))
- Add Tequilapi client timeout [\#460](https://github.com/mysteriumnetwork/node/pull/460) ([tcharding](https://github.com/tcharding))
- redirect stderr to stdout for goimports \#450 [\#451](https://github.com/mysteriumnetwork/node/pull/451) ([u5surf](https://github.com/u5surf))
- Dummy connection which does Noop tunnel [\#446](https://github.com/mysteriumnetwork/node/pull/446) ([Waldz](https://github.com/Waldz))

## [0.4.0](https://github.com/mysteriumnetwork/node/tree/0.4.0) (2018-10-18)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.3.2...0.4.0)

**Implemented enhancements:**

- Added --experiment-promise-check flag [\#464](https://github.com/mysteriumnetwork/node/pull/464) ([soffokl](https://github.com/soffokl))

**Fixed bugs:**

- Goimports check do nothing since goimport is not installed for CI [\#450](https://github.com/mysteriumnetwork/node/issues/450)
- Fixed getting version from tag [\#467](https://github.com/mysteriumnetwork/node/pull/467) ([soffokl](https://github.com/soffokl))

**Closed issues:**

- Payments. Provider stores Consumer promises \(state\) [\#347](https://github.com/mysteriumnetwork/node/issues/347)

**Merged pull requests:**

- Update development environment setup docs [\#461](https://github.com/mysteriumnetwork/node/pull/461) ([tcharding](https://github.com/tcharding))
- Force bin/build to run [\#459](https://github.com/mysteriumnetwork/node/pull/459) ([tcharding](https://github.com/tcharding))
- Reduce noise when running `check\_golint` [\#457](https://github.com/mysteriumnetwork/node/pull/457) ([tcharding](https://github.com/tcharding))
- Check we have `goimports` [\#453](https://github.com/mysteriumnetwork/node/pull/453) ([tcharding](https://github.com/tcharding))
- Run goimports with builder [\#452](https://github.com/mysteriumnetwork/node/pull/452) ([tadovas](https://github.com/tadovas))
- Cherrypick of "Identity check disabled by default, added flag to enable it. \#363" [\#445](https://github.com/mysteriumnetwork/node/pull/445) ([Waldz](https://github.com/Waldz))
- Added fetchQuality to the proposals endpoint with a fake data [\#444](https://github.com/mysteriumnetwork/node/pull/444) ([soffokl](https://github.com/soffokl))
- Added promises validation and storing for the service provider. \#347 [\#404](https://github.com/mysteriumnetwork/node/pull/404) ([soffokl](https://github.com/soffokl))
- Move towards a more modular connection manager [\#385](https://github.com/mysteriumnetwork/node/pull/385) ([vkuznecovas](https://github.com/vkuznecovas))
- Compile in geoip2 database and avoid external file reading [\#281](https://github.com/mysteriumnetwork/node/pull/281) ([tadovas](https://github.com/tadovas))

## [0.3.2](https://github.com/mysteriumnetwork/node/tree/0.3.2) (2018-10-15)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.3.1...0.3.2)

**Implemented enhancements:**

- Mark PR as stale in 2 weeks and close it in 2 more weeks [\#438](https://github.com/mysteriumnetwork/node/pull/438) ([soffokl](https://github.com/soffokl))
- Push master to dev pre-release [\#423](https://github.com/mysteriumnetwork/node/pull/423) ([soffokl](https://github.com/soffokl))
- Cli changes [\#421](https://github.com/mysteriumnetwork/node/pull/421) ([soffokl](https://github.com/soffokl))

**Fixed bugs:**

- 0.3.0 standalone has wrong 'config' directory structure [\#415](https://github.com/mysteriumnetwork/node/issues/415)
- Cherry pick of "Added back agreed-terms-and-conditions flag" [\#442](https://github.com/mysteriumnetwork/node/pull/442) ([soffokl](https://github.com/soffokl))
- Fixed build version for deb packages [\#430](https://github.com/mysteriumnetwork/node/pull/430) ([soffokl](https://github.com/soffokl))
- Fixed travis stages [\#427](https://github.com/mysteriumnetwork/node/pull/427) ([soffokl](https://github.com/soffokl))

**Closed issues:**

- Push 'dev' releases to public storage [\#416](https://github.com/mysteriumnetwork/node/issues/416)
- move away from satori/go.uuid [\#376](https://github.com/mysteriumnetwork/node/issues/376)
- Release 0.3 version with Paid ID [\#363](https://github.com/mysteriumnetwork/node/issues/363)
- Payments. Provider constantly notifies current Consumer's balance [\#346](https://github.com/mysteriumnetwork/node/issues/346)
- Payments. Consumer issues new promise [\#345](https://github.com/mysteriumnetwork/node/issues/345)
- Change --cli argument to "myst cli" subcommand [\#335](https://github.com/mysteriumnetwork/node/issues/335)

**Merged pull requests:**

- Fix broken build card url [\#448](https://github.com/mysteriumnetwork/node/pull/448) ([tadovas](https://github.com/tadovas))
- moved to gofrs/uuid instead of satori/uuid [\#422](https://github.com/mysteriumnetwork/node/pull/422) ([vkuznecovas](https://github.com/vkuznecovas))
- Promise balance notifications [\#381](https://github.com/mysteriumnetwork/node/pull/381) ([Waldz](https://github.com/Waldz))
- Consumer issues new promise for the provider [\#362](https://github.com/mysteriumnetwork/node/pull/362) ([soffokl](https://github.com/soffokl))

## [0.3.1](https://github.com/mysteriumnetwork/node/tree/0.3.1) (2018-10-02)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.3.0...0.3.1)

**Implemented enhancements:**

- Smaller final binaries [\#410](https://github.com/mysteriumnetwork/node/pull/410) ([soffokl](https://github.com/soffokl))
- Removed building of client and server, only myst left [\#406](https://github.com/mysteriumnetwork/node/pull/406) ([soffokl](https://github.com/soffokl))

**Fixed bugs:**

- Fix travis status badge  [\#397](https://github.com/mysteriumnetwork/node/issues/397)
- myst deb packages does not provide service after installation [\#395](https://github.com/mysteriumnetwork/node/issues/395)
- Backport of configuration fixes from \#406 [\#418](https://github.com/mysteriumnetwork/node/pull/418) ([soffokl](https://github.com/soffokl))
- Fixed service starting in the dep packages [\#413](https://github.com/mysteriumnetwork/node/pull/413) ([soffokl](https://github.com/soffokl))
- Updated Readme.md to fit latest release and iptables rule fix backport [\#393](https://github.com/mysteriumnetwork/node/pull/393) ([soffokl](https://github.com/soffokl))

**Closed issues:**

- Setup automatic stale issue actions with comments [\#359](https://github.com/mysteriumnetwork/node/issues/359)
- SPIKE kill switch [\#354](https://github.com/mysteriumnetwork/node/issues/354)
- Freeze stable releases from branch "release/0.1" [\#343](https://github.com/mysteriumnetwork/node/issues/343)
- Move commit history to repo "go-openvpn" [\#342](https://github.com/mysteriumnetwork/node/issues/342)
- Decouple "openvpn" package from Mysterium logic [\#341](https://github.com/mysteriumnetwork/node/issues/341)
- Make openvpn tunnel setup mandatory for process [\#340](https://github.com/mysteriumnetwork/node/issues/340)
- Discovery doest not work with RPC witch has POA mining [\#337](https://github.com/mysteriumnetwork/node/issues/337)
- E2E tests. Point localnet new Discovery version [\#336](https://github.com/mysteriumnetwork/node/issues/336)
- Server response invalid: 400 BAD REQUEST [\#125](https://github.com/mysteriumnetwork/node/issues/125)

**Merged pull requests:**

- Added stale configuration for auto closing inactive issues. [\#412](https://github.com/mysteriumnetwork/node/pull/412) ([soffokl](https://github.com/soffokl))
- Updated golang to 1.11 [\#409](https://github.com/mysteriumnetwork/node/pull/409) ([soffokl](https://github.com/soffokl))
- Dont expose session transport DTOs to external world [\#405](https://github.com/mysteriumnetwork/node/pull/405) ([Waldz](https://github.com/Waldz))
- CI improvements [\#403](https://github.com/mysteriumnetwork/node/pull/403) ([Waldz](https://github.com/Waldz))
- updated travis badge url [\#400](https://github.com/mysteriumnetwork/node/pull/400) ([vkuznecovas](https://github.com/vkuznecovas))
- added a step to update go report cards on master commit [\#396](https://github.com/mysteriumnetwork/node/pull/396) ([vkuznecovas](https://github.com/vkuznecovas))
- Inject Openvpn service on service provider side [\#392](https://github.com/mysteriumnetwork/node/pull/392) ([Waldz](https://github.com/Waldz))

## [0.3.0](https://github.com/mysteriumnetwork/node/tree/0.3.0) (2018-09-20)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.2.3...0.3.0)

**Implemented enhancements:**

- Docker host requires additional iptables NAT rule [\#16](https://github.com/mysteriumnetwork/node/issues/16)
- Node supports several clients at same time  [\#13](https://github.com/mysteriumnetwork/node/issues/13)
- mysterium\_client Windows support [\#3](https://github.com/mysteriumnetwork/node/issues/3)
- Decouple session storage [\#380](https://github.com/mysteriumnetwork/node/pull/380) ([Waldz](https://github.com/Waldz))
- Custom cancelable implementation replaced with context [\#378](https://github.com/mysteriumnetwork/node/pull/378) ([soffokl](https://github.com/soffokl))
- CI builds for all artifacts with a single client+server binary [\#327](https://github.com/mysteriumnetwork/node/pull/327) ([soffokl](https://github.com/soffokl))
- MYST-689. Added sending last stats on disconnect [\#325](https://github.com/mysteriumnetwork/node/pull/325) ([soffokl](https://github.com/soffokl))
- MYST-304. Enabled data race detector for unit tests [\#322](https://github.com/mysteriumnetwork/node/pull/322) ([soffokl](https://github.com/soffokl))
- MYST-178. Fixed import path in packages [\#321](https://github.com/mysteriumnetwork/node/pull/321) ([soffokl](https://github.com/soffokl))
- Migration from glide to dep [\#319](https://github.com/mysteriumnetwork/node/pull/319) ([soffokl](https://github.com/soffokl))

**Fixed bugs:**

- Wrong Go package names [\#14](https://github.com/mysteriumnetwork/node/issues/14)
- Node supports several clients at same time  [\#13](https://github.com/mysteriumnetwork/node/issues/13)
- ARM64 build for node [\#11](https://github.com/mysteriumnetwork/node/issues/11)
- Ubuntu 16.04 Node Failing to Start [\#8](https://github.com/mysteriumnetwork/node/issues/8)
- Fixed of building docker images for Ubuntu [\#328](https://github.com/mysteriumnetwork/node/pull/328) ([soffokl](https://github.com/soffokl))

**Closed issues:**

- Move "openvpn" package to separate repository [\#344](https://github.com/mysteriumnetwork/node/issues/344)
- Lots of spelling issues with online scanning [\#308](https://github.com/mysteriumnetwork/node/issues/308)
- Installation Link & Download [\#243](https://github.com/mysteriumnetwork/node/issues/243)
- How to use client in Maxos? [\#24](https://github.com/mysteriumnetwork/node/issues/24)

**Merged pull requests:**

- Fixed upload artifacts to release and set iptables rules [\#388](https://github.com/mysteriumnetwork/node/pull/388) ([soffokl](https://github.com/soffokl))
- Fixed uploading release images [\#384](https://github.com/mysteriumnetwork/node/pull/384) ([soffokl](https://github.com/soffokl))
- Identity check disabled by default, added flag to enable it. \#363 [\#383](https://github.com/mysteriumnetwork/node/pull/383) ([soffokl](https://github.com/soffokl))
- Move to go openvpn [\#379](https://github.com/mysteriumnetwork/node/pull/379) ([vkuznecovas](https://github.com/vkuznecovas))
- Remove openvpn session manager [\#377](https://github.com/mysteriumnetwork/node/pull/377) ([Waldz](https://github.com/Waldz))
- Decouple openvpn [\#371](https://github.com/mysteriumnetwork/node/pull/371) ([vkuznecovas](https://github.com/vkuznecovas))
- E2E use discovery version with registered id [\#370](https://github.com/mysteriumnetwork/node/pull/370) ([interro](https://github.com/interro))
- Updated README.md [\#356](https://github.com/mysteriumnetwork/node/pull/356) ([ignasbernotas](https://github.com/ignasbernotas))
- Pushing myst image to both myst and mysterium-node repos for the migration period [\#332](https://github.com/mysteriumnetwork/node/pull/332) ([soffokl](https://github.com/soffokl))
- MYST-726 Make openvpn tunnel setup mandatory for process [\#331](https://github.com/mysteriumnetwork/node/pull/331) ([Waldz](https://github.com/Waldz))
- Feature/myst 708 kill switch [\#330](https://github.com/mysteriumnetwork/node/pull/330) ([zolia](https://github.com/zolia))
- MYST-703 Refactor CLI flag bootstraping [\#329](https://github.com/mysteriumnetwork/node/pull/329) ([Waldz](https://github.com/Waldz))
- Fixed some misspells. \#308 [\#326](https://github.com/mysteriumnetwork/node/pull/326) ([soffokl](https://github.com/soffokl))
- MYST-682 Move mysterium\_server logic to "mysterium\_node service" subcommand [\#324](https://github.com/mysteriumnetwork/node/pull/324) ([Waldz](https://github.com/Waldz))
- Feature/myst 684 e2e tests registers identity [\#323](https://github.com/mysteriumnetwork/node/pull/323) ([zolia](https://github.com/zolia))
- Start using goimports [\#317](https://github.com/mysteriumnetwork/node/pull/317) ([Waldz](https://github.com/Waldz))
- MYST-688 Start having subcommands of "mysterium\_node" [\#316](https://github.com/mysteriumnetwork/node/pull/316) ([Waldz](https://github.com/Waldz))
- Upload artifacts to S3 moved to HTTPS [\#315](https://github.com/mysteriumnetwork/node/pull/315) ([soffokl](https://github.com/soffokl))
- MYST-681 Start having command "mysterium\_node" [\#314](https://github.com/mysteriumnetwork/node/pull/314) ([Waldz](https://github.com/Waldz))
- Remove @donce from code owners [\#313](https://github.com/mysteriumnetwork/node/pull/313) ([donce](https://github.com/donce))
- Updated go-homedir [\#312](https://github.com/mysteriumnetwork/node/pull/312) ([ignasbernotas](https://github.com/ignasbernotas))
- Feature/myst-572 node rejects unregistered consumers [\#311](https://github.com/mysteriumnetwork/node/pull/311) ([tadovas](https://github.com/tadovas))
- Feature/myst 586 endpoint with identity registration status [\#310](https://github.com/mysteriumnetwork/node/pull/310) ([tadovas](https://github.com/tadovas))
- Trigger packaging on master branch commits and tags [\#309](https://github.com/mysteriumnetwork/node/pull/309) ([tadovas](https://github.com/tadovas))
- Run packaging checks only on PR [\#307](https://github.com/mysteriumnetwork/node/pull/307) ([Waldz](https://github.com/Waldz))
- Pick fixes from 0.2 [\#306](https://github.com/mysteriumnetwork/node/pull/306) ([Waldz](https://github.com/Waldz))
- MYST-643 Make CI parallel building [\#287](https://github.com/mysteriumnetwork/node/pull/287) ([tadovas](https://github.com/tadovas))
- MYST-445 Make install and readme instructions great again [\#284](https://github.com/mysteriumnetwork/node/pull/284) ([tadovas](https://github.com/tadovas))

## [0.2.3](https://github.com/mysteriumnetwork/node/tree/0.2.3) (2018-07-25)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.2.2...0.2.3)

**Closed issues:**

- Can I run a node from the US Virgin Islands? [\#300](https://github.com/mysteriumnetwork/node/issues/300)

**Merged pull requests:**

- Pack binaries with .exe for Windows to release [\#305](https://github.com/mysteriumnetwork/node/pull/305) ([interro](https://github.com/interro))
- Improvement/Merge client and server network options into single structure [\#303](https://github.com/mysteriumnetwork/node/pull/303) ([tadovas](https://github.com/tadovas))
- Better errors handling [\#302](https://github.com/mysteriumnetwork/node/pull/302) ([soffokl](https://github.com/soffokl))
- Fixed typo in the CONTRIBUTING guide [\#301](https://github.com/mysteriumnetwork/node/pull/301) ([soffokl](https://github.com/soffokl))
- darwin nat implementation [\#298](https://github.com/mysteriumnetwork/node/pull/298) ([zolia](https://github.com/zolia))

## [0.2.2](https://github.com/mysteriumnetwork/node/tree/0.2.2) (2018-07-17)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.2.1...0.2.2)

**Merged pull requests:**

- HOTFIX Export build arguments to global scope [\#299](https://github.com/mysteriumnetwork/node/pull/299) ([Waldz](https://github.com/Waldz))

## [0.2.1](https://github.com/mysteriumnetwork/node/tree/0.2.1) (2018-07-13)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.2.0...0.2.1)

**Merged pull requests:**

- HOTFIX: copy proper config files [\#296](https://github.com/mysteriumnetwork/node/pull/296) ([Waldz](https://github.com/Waldz))
- verbose copy [\#294](https://github.com/mysteriumnetwork/node/pull/294) ([zolia](https://github.com/zolia))
- copy client configs travis way [\#293](https://github.com/mysteriumnetwork/node/pull/293) ([zolia](https://github.com/zolia))
- Improvement/naive detection of nat [\#292](https://github.com/mysteriumnetwork/node/pull/292) ([tadovas](https://github.com/tadovas))
- HOTFIX: copy proper config files  [\#291](https://github.com/mysteriumnetwork/node/pull/291) ([Waldz](https://github.com/Waldz))
- HOTFIX: copy proper config files [\#290](https://github.com/mysteriumnetwork/node/pull/290) ([zolia](https://github.com/zolia))

## [0.2.0](https://github.com/mysteriumnetwork/node/tree/0.2.0) (2018-07-12)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.0.6...0.2.0)

**Fixed bugs:**

- Fix a loop error [\#20](https://github.com/mysteriumnetwork/node/issues/20)
- Fix syntax error in Dockerfile [\#9](https://github.com/mysteriumnetwork/node/issues/9)
- Remove the upper cap dependence on openvpn version \<\<2.3.3 [\#5](https://github.com/mysteriumnetwork/node/issues/5)

**Closed issues:**

- Use govendor to manage dependent packages [\#27](https://github.com/mysteriumnetwork/node/issues/27)
- Add a Makefile for convenience [\#23](https://github.com/mysteriumnetwork/node/issues/23)
- Traffic statistics "Data transfered" for node sessions [\#4](https://github.com/mysteriumnetwork/node/issues/4)

**Merged pull requests:**

- HOTFIX Fix ternary expression \(with default value\) in Bash scripts [\#289](https://github.com/mysteriumnetwork/node/pull/289) ([Waldz](https://github.com/Waldz))
- add exec perm for update dns script [\#288](https://github.com/mysteriumnetwork/node/pull/288) ([zolia](https://github.com/zolia))
- MYST-637 Fix cancelable test deadlock [\#286](https://github.com/mysteriumnetwork/node/pull/286) ([tadovas](https://github.com/tadovas))
- MYST-577 Fix missing version during "dev" builds [\#285](https://github.com/mysteriumnetwork/node/pull/285) ([Waldz](https://github.com/Waldz))
- MYST-564 Fix config directory in standalone artifacts [\#283](https://github.com/mysteriumnetwork/node/pull/283) ([Waldz](https://github.com/Waldz))
- Add 'healthcheck' command to CLI [\#282](https://github.com/mysteriumnetwork/node/pull/282) ([donce](https://github.com/donce))
- MYST-566 Release Docker artifacts  [\#280](https://github.com/mysteriumnetwork/node/pull/280) ([Waldz](https://github.com/Waldz))
- Lower timeout of unit tests from default 10m to 3 m [\#279](https://github.com/mysteriumnetwork/node/pull/279) ([tadovas](https://github.com/tadovas))
- feature/MYST-596 DNS update for windows [\#278](https://github.com/mysteriumnetwork/node/pull/278) ([tadovas](https://github.com/tadovas))
- MYST-590 terms for deb install [\#277](https://github.com/mysteriumnetwork/node/pull/277) ([zolia](https://github.com/zolia))
- MYST-565 Release debian artefacts [\#276](https://github.com/mysteriumnetwork/node/pull/276) ([Waldz](https://github.com/Waldz))
- Fix broken linking to metadata package of build info [\#275](https://github.com/mysteriumnetwork/node/pull/275) ([tadovas](https://github.com/tadovas))
- MYST-614 Stop openvpn executable in OS portable way [\#273](https://github.com/mysteriumnetwork/node/pull/273) ([tadovas](https://github.com/tadovas))
- MYST-569 get rid of build time env variables [\#272](https://github.com/mysteriumnetwork/node/pull/272) ([tadovas](https://github.com/tadovas))
- Document all API endpoints [\#271](https://github.com/mysteriumnetwork/node/pull/271) ([interro](https://github.com/interro))
- Fix response headers on tequilapi responses [\#269](https://github.com/mysteriumnetwork/node/pull/269) ([tadovas](https://github.com/tadovas))
- MYST-563 Compile version info executable [\#268](https://github.com/mysteriumnetwork/node/pull/268) ([Waldz](https://github.com/Waldz))
- MYST-559 - fix forever hanging openvpn client if process exits too early [\#267](https://github.com/mysteriumnetwork/node/pull/267) ([tadovas](https://github.com/tadovas))
- Playground with go swagger [\#265](https://github.com/mysteriumnetwork/node/pull/265) ([interro](https://github.com/interro))
- MYST-578 Start using XGO in CI flow [\#264](https://github.com/mysteriumnetwork/node/pull/264) ([Waldz](https://github.com/Waldz))
- Feature/myst 437 sudoers for debs [\#263](https://github.com/mysteriumnetwork/node/pull/263) ([zolia](https://github.com/zolia))
- MYST-549 Terms & conditions [\#262](https://github.com/mysteriumnetwork/node/pull/262) ([donce](https://github.com/donce))
- HOTFIX Connect all checks to CI flow [\#260](https://github.com/mysteriumnetwork/node/pull/260) ([Waldz](https://github.com/Waldz))
- Finally travis fixed travis-ci/travis-ci\#9312 [\#259](https://github.com/mysteriumnetwork/node/pull/259) ([tadovas](https://github.com/tadovas))
- MYST-546/Refactor spaces in option values handling [\#258](https://github.com/mysteriumnetwork/node/pull/258) ([tadovas](https://github.com/tadovas))
- MYST-495 Show license on startup with subcommands for more info [\#257](https://github.com/mysteriumnetwork/node/pull/257) ([donce](https://github.com/donce))
- MYST-491 Add CI check for copyright headers [\#256](https://github.com/mysteriumnetwork/node/pull/256) ([donce](https://github.com/donce))
- MYS-28 Added win/osx/linux cross-compilation via xgo [\#255](https://github.com/mysteriumnetwork/node/pull/255) ([Waldz](https://github.com/Waldz))
- README, INSTALL and CONTRIBUTING files updated [\#254](https://github.com/mysteriumnetwork/node/pull/254) ([zolia](https://github.com/zolia))
- Cleanup terminal colors after printing tests output [\#253](https://github.com/mysteriumnetwork/node/pull/253) ([donce](https://github.com/donce))
- Moved openvpn management from unix socket to TCP to support windows [\#252](https://github.com/mysteriumnetwork/node/pull/252) ([ignasbernotas](https://github.com/ignasbernotas))
- generic request timeout [\#251](https://github.com/mysteriumnetwork/node/pull/251) ([zolia](https://github.com/zolia))
- Add go fmt check to CI [\#249](https://github.com/mysteriumnetwork/node/pull/249) ([donce](https://github.com/donce))
- Feature/myst 497 improve session authorisation [\#248](https://github.com/mysteriumnetwork/node/pull/248) ([zolia](https://github.com/zolia))
- MYST-490 license header [\#245](https://github.com/mysteriumnetwork/node/pull/245) ([donce](https://github.com/donce))
- Add license headers [\#242](https://github.com/mysteriumnetwork/node/pull/242) ([donce](https://github.com/donce))
- Do not release same version twice when tagging semantic version [\#240](https://github.com/mysteriumnetwork/node/pull/240) ([donce](https://github.com/donce))
- Hotfix/unregistered provider print fix [\#238](https://github.com/mysteriumnetwork/node/pull/238) ([zolia](https://github.com/zolia))
- Add go vet to check for dead code [\#237](https://github.com/mysteriumnetwork/node/pull/237) ([donce](https://github.com/donce))
- Feature/myst-470 pass vpn options structure [\#236](https://github.com/mysteriumnetwork/node/pull/236) ([tadovas](https://github.com/tadovas))
- MYST 532 Tequilapi new Location Endpoint [\#235](https://github.com/mysteriumnetwork/node/pull/235) ([interro](https://github.com/interro))
- Add explicit exit notify option only if UDP protocol is specified [\#234](https://github.com/mysteriumnetwork/node/pull/234) ([tadovas](https://github.com/tadovas))
- Add notify on explicit exit openvpn flag [\#233](https://github.com/mysteriumnetwork/node/pull/233) ([tadovas](https://github.com/tadovas))
- Feature/myst 372 unregister proposal [\#232](https://github.com/mysteriumnetwork/node/pull/232) ([zolia](https://github.com/zolia))
- MYST-511 / Add http response header configurable timeout and test [\#231](https://github.com/mysteriumnetwork/node/pull/231) ([tadovas](https://github.com/tadovas))
- retry register proposal [\#230](https://github.com/mysteriumnetwork/node/pull/230) ([zolia](https://github.com/zolia))
- Update CONTRIBUTING doc [\#229](https://github.com/mysteriumnetwork/node/pull/229) ([zyfdegh](https://github.com/zyfdegh))
- Refactor connect to return custom error [\#228](https://github.com/mysteriumnetwork/node/pull/228) ([donce](https://github.com/donce))
- Add codeowners file [\#227](https://github.com/mysteriumnetwork/node/pull/227) ([tadovas](https://github.com/tadovas))
- HOTFIX Unify script names [\#223](https://github.com/mysteriumnetwork/node/pull/223) ([Waldz](https://github.com/Waldz))
- MYST-498 Cancel connection returns 499 [\#222](https://github.com/mysteriumnetwork/node/pull/222) ([donce](https://github.com/donce))
- fix reconnect issue to gnatsd [\#221](https://github.com/mysteriumnetwork/node/pull/221) ([zolia](https://github.com/zolia))
- Wait for database to become ready before performing migration [\#220](https://github.com/mysteriumnetwork/node/pull/220) ([tadovas](https://github.com/tadovas))
- Stop tequila after current request is finished, instead of delay [\#219](https://github.com/mysteriumnetwork/node/pull/219) ([tadovas](https://github.com/tadovas))
- Refactor to fix some linter errors [\#218](https://github.com/mysteriumnetwork/node/pull/218) ([donce](https://github.com/donce))
- Disable caching on all tequilapi responses [\#217](https://github.com/mysteriumnetwork/node/pull/217) ([tadovas](https://github.com/tadovas))
- Add custom status code for cancelled connection [\#216](https://github.com/mysteriumnetwork/node/pull/216) ([tadovas](https://github.com/tadovas))
- hotfix/fix-non-deterministic-cancelable-tests [\#215](https://github.com/mysteriumnetwork/node/pull/215) ([tadovas](https://github.com/tadovas))
- HOTFIX. Fix date of Copyright [\#214](https://github.com/mysteriumnetwork/node/pull/214) ([Waldz](https://github.com/Waldz))
- Feature/MYST-431-make-connect-method-cancelable [\#212](https://github.com/mysteriumnetwork/node/pull/212) ([tadovas](https://github.com/tadovas))
- e2e tests use new discovery api release [\#211](https://github.com/mysteriumnetwork/node/pull/211) ([interro](https://github.com/interro))
- removed privileged and mknod cap; docker entrypoints reworked [\#210](https://github.com/mysteriumnetwork/node/pull/210) ([zolia](https://github.com/zolia))
- Improvement/discard env variables in alpine images [\#209](https://github.com/mysteriumnetwork/node/pull/209) ([tadovas](https://github.com/tadovas))
- new server option to specify protocol [\#208](https://github.com/mysteriumnetwork/node/pull/208) ([zolia](https://github.com/zolia))
- 345 Client Get Country [\#207](https://github.com/mysteriumnetwork/node/pull/207) ([interro](https://github.com/interro))
- MYST-420 / manager.Connect now sets itself into Connecting state [\#206](https://github.com/mysteriumnetwork/node/pull/206) ([ignasbernotas](https://github.com/ignasbernotas))
- Use 127.0.0.1 by default for tequilapi to make CLI work with daemon client [\#205](https://github.com/mysteriumnetwork/node/pull/205) ([donce](https://github.com/donce))
- Feature/implement e2e client connects to node test [\#204](https://github.com/mysteriumnetwork/node/pull/204) ([tadovas](https://github.com/tadovas))
- Hotfix/enable slack notifications [\#203](https://github.com/mysteriumnetwork/node/pull/203) ([tadovas](https://github.com/tadovas))
- MYST-413 Optimize dockers to have more cached layers [\#202](https://github.com/mysteriumnetwork/node/pull/202) ([Waldz](https://github.com/Waldz))
- Fix travis yml for travis plugin downgrade [\#201](https://github.com/mysteriumnetwork/node/pull/201) ([tadovas](https://github.com/tadovas))
- Downgrade travis pages provider, to avoid issue with release 1.9.0 [\#200](https://github.com/mysteriumnetwork/node/pull/200) ([tadovas](https://github.com/tadovas))
- Feature/add-ipify-runtime-option [\#199](https://github.com/mysteriumnetwork/node/pull/199) ([tadovas](https://github.com/tadovas))
- docker builder script; package\_all remade to acomodate new builder [\#198](https://github.com/mysteriumnetwork/node/pull/198) ([zolia](https://github.com/zolia))
- MYST-413 Fix broken Alpine image [\#197](https://github.com/mysteriumnetwork/node/pull/197) ([Waldz](https://github.com/Waldz))
- Fix failing test when running on Go 1.10 [\#196](https://github.com/mysteriumnetwork/node/pull/196) ([donce](https://github.com/donce))
- Feature/enable-e2e-tests-on-pull-requests [\#195](https://github.com/mysteriumnetwork/node/pull/195) ([tadovas](https://github.com/tadovas))
- Generate test files in testdataouput dir [\#193](https://github.com/mysteriumnetwork/node/pull/193) ([tadovas](https://github.com/tadovas))
- build info in struct /healthcheck [\#192](https://github.com/mysteriumnetwork/node/pull/192) ([shroomist](https://github.com/shroomist))
- Bugfix/MYST-393 [\#191](https://github.com/mysteriumnetwork/node/pull/191) ([tadovas](https://github.com/tadovas))
- Make log timestamps readable \(UTC timezone\) [\#190](https://github.com/mysteriumnetwork/node/pull/190) ([tadovas](https://github.com/tadovas))
- feature/MYST-320 Expose client build info on tequila api [\#189](https://github.com/mysteriumnetwork/node/pull/189) ([tadovas](https://github.com/tadovas))
- feature/MYST-387 [\#188](https://github.com/mysteriumnetwork/node/pull/188) ([tadovas](https://github.com/tadovas))
- Fail fast if could not locate openvpn binary on both client and node [\#187](https://github.com/mysteriumnetwork/node/pull/187) ([tadovas](https://github.com/tadovas))
- Hotfix/myst 361 pass api nats ips via cmdline [\#186](https://github.com/mysteriumnetwork/node/pull/186) ([zolia](https://github.com/zolia))
- Revert ipify timeout to 1 minute [\#185](https://github.com/mysteriumnetwork/node/pull/185) ([donce](https://github.com/donce))
- Stop session duration when vpn client exits [\#183](https://github.com/mysteriumnetwork/node/pull/183) ([donce](https://github.com/donce))
- Fix interval for stats sender [\#182](https://github.com/mysteriumnetwork/node/pull/182) ([donce](https://github.com/donce))
- Disable connection caching for http clients in both node and client [\#181](https://github.com/mysteriumnetwork/node/pull/181) ([tadovas](https://github.com/tadovas))
- MYST-323 Increase stats update frequency [\#180](https://github.com/mysteriumnetwork/node/pull/180) ([donce](https://github.com/donce))
- MYST-362 prevent server premature shutdown on missing discovery ping [\#179](https://github.com/mysteriumnetwork/node/pull/179) ([zolia](https://github.com/zolia))
- Add simple integration test for client with CLI [\#178](https://github.com/mysteriumnetwork/node/pull/178) ([donce](https://github.com/donce))
- Remove temp archive dir [\#177](https://github.com/mysteriumnetwork/node/pull/177) ([tadovas](https://github.com/tadovas))
- Improvement/archive client server as tar gz [\#176](https://github.com/mysteriumnetwork/node/pull/176) ([tadovas](https://github.com/tadovas))
- Add some missing comments to reduce linter warnings [\#175](https://github.com/mysteriumnetwork/node/pull/175) ([donce](https://github.com/donce))
- Increase NATS sender timeout to avoid early timeouts [\#174](https://github.com/mysteriumnetwork/node/pull/174) ([donce](https://github.com/donce))
- feature/MYST-149 command line option to specify openvpn executable [\#173](https://github.com/mysteriumnetwork/node/pull/173) ([tadovas](https://github.com/tadovas))
- MYST 361: add -discovery-address and -broker-address command line options [\#172](https://github.com/mysteriumnetwork/node/pull/172) ([zolia](https://github.com/zolia))
- Hotfix/enable history preserve for pages provider [\#171](https://github.com/mysteriumnetwork/node/pull/171) ([tadovas](https://github.com/tadovas))
- Improvement/build artifacts for macos and linux [\#170](https://github.com/mysteriumnetwork/node/pull/170) ([tadovas](https://github.com/tadovas))
- Improvement/upload mysterium binaries on master branch commit [\#169](https://github.com/mysteriumnetwork/node/pull/169) ([tadovas](https://github.com/tadovas))
- Fix linter warnings [\#168](https://github.com/mysteriumnetwork/node/pull/168) ([donce](https://github.com/donce))
- MYST 355: limit openvpn client reconnects [\#167](https://github.com/mysteriumnetwork/node/pull/167) ([zolia](https://github.com/zolia))
- MYST-324 Return "Service Unavailable" when ipify is unavailable [\#166](https://github.com/mysteriumnetwork/node/pull/166) ([donce](https://github.com/donce))
- HOTFIX Dont omit location database on --location.database="" [\#165](https://github.com/mysteriumnetwork/node/pull/165) ([Waldz](https://github.com/Waldz))
- Use new discovery api endpoint urls [\#164](https://github.com/mysteriumnetwork/node/pull/164) ([donce](https://github.com/donce))
- HOTFIX. Freeze Openvpn dependency [\#162](https://github.com/mysteriumnetwork/node/pull/162) ([Waldz](https://github.com/Waldz))
- Include providerID to statistics sent by client [\#161](https://github.com/mysteriumnetwork/node/pull/161) ([interro](https://github.com/interro))
- MYST-333 Fix node auth error when client ids are several number length [\#160](https://github.com/mysteriumnetwork/node/pull/160) ([tadovas](https://github.com/tadovas))
- Fix failing deb package script, when script is launched as standalone [\#159](https://github.com/mysteriumnetwork/node/pull/159) ([tadovas](https://github.com/tadovas))
- New cli command - stop, stops mysterium client [\#158](https://github.com/mysteriumnetwork/node/pull/158) ([tadovas](https://github.com/tadovas))
- Feature/myst-297 service provider openvpn state reporting [\#157](https://github.com/mysteriumnetwork/node/pull/157) ([tadovas](https://github.com/tadovas))
- Add signal handler for mysterium service to stop gracefully [\#156](https://github.com/mysteriumnetwork/node/pull/156) ([tadovas](https://github.com/tadovas))
- Move `client\_promise` to `client/promise` package [\#155](https://github.com/mysteriumnetwork/node/pull/155) ([donce](https://github.com/donce))
- Rename client\_connection package name to Go naming conventions [\#154](https://github.com/mysteriumnetwork/node/pull/154) ([donce](https://github.com/donce))
- MYST-298 Possibility to inject identity to Docker container [\#153](https://github.com/mysteriumnetwork/node/pull/153) ([Waldz](https://github.com/Waldz))
- Add openvpn error creation handling [\#152](https://github.com/mysteriumnetwork/node/pull/152) ([tadovas](https://github.com/tadovas))
- Improvement/add caching layer [\#151](https://github.com/mysteriumnetwork/node/pull/151) ([tadovas](https://github.com/tadovas))
- Feature/myst 79 regen openvpn keys [\#150](https://github.com/mysteriumnetwork/node/pull/150) ([zolia](https://github.com/zolia))
- Add build widget to project description [\#149](https://github.com/mysteriumnetwork/node/pull/149) ([tadovas](https://github.com/tadovas))
- Travis file added [\#148](https://github.com/mysteriumnetwork/node/pull/148) ([tadovas](https://github.com/tadovas))
- Feature/myst-227 openvpn client async state changes [\#147](https://github.com/mysteriumnetwork/node/pull/147) ([tadovas](https://github.com/tadovas))
- MYST-300 Update client's system dns [\#146](https://github.com/mysteriumnetwork/node/pull/146) ([Waldz](https://github.com/Waldz))
- Split CLI command [\#144](https://github.com/mysteriumnetwork/node/pull/144) ([donce](https://github.com/donce))
- Add Go report card to README.md [\#143](https://github.com/mysteriumnetwork/node/pull/143) ([donce](https://github.com/donce))
- Remove mysterium fake command [\#142](https://github.com/mysteriumnetwork/node/pull/142) ([donce](https://github.com/donce))
- MYST-300 Update client's system dns [\#141](https://github.com/mysteriumnetwork/node/pull/141) ([Waldz](https://github.com/Waldz))
- Remove mysterium monitor command [\#140](https://github.com/mysteriumnetwork/node/pull/140) ([donce](https://github.com/donce))
- MYST-291 Karakiri endpoint [\#139](https://github.com/mysteriumnetwork/node/pull/139) ([donce](https://github.com/donce))
- HOSTFIX fixed networksetup path [\#138](https://github.com/mysteriumnetwork/node/pull/138) ([Waldz](https://github.com/Waldz))
- MYST-298 Rename nodeKey to providerId [\#137](https://github.com/mysteriumnetwork/node/pull/137) ([Waldz](https://github.com/Waldz))
- execute script to update mac system dns servers after openvpn connect… [\#136](https://github.com/mysteriumnetwork/node/pull/136) ([zolia](https://github.com/zolia))
- HOTFIX One step towards single node [\#135](https://github.com/mysteriumnetwork/node/pull/135) ([Waldz](https://github.com/Waldz))
- Cors preflight request detection improved [\#134](https://github.com/mysteriumnetwork/node/pull/134) ([tadovas](https://github.com/tadovas))
- HOTFIX Build correct docker file during release [\#133](https://github.com/mysteriumnetwork/node/pull/133) ([Waldz](https://github.com/Waldz))
- Rename 'service\_discovery' package to 'discovery' [\#132](https://github.com/mysteriumnetwork/node/pull/132) ([donce](https://github.com/donce))
- HOTFIX One step towards single node [\#131](https://github.com/mysteriumnetwork/node/pull/131) ([Waldz](https://github.com/Waldz))
- HOTFIX Usecase of undetectable country [\#130](https://github.com/mysteriumnetwork/node/pull/130) ([Waldz](https://github.com/Waldz))
- Use default min-confidence for linter on new code [\#129](https://github.com/mysteriumnetwork/node/pull/129) ([donce](https://github.com/donce))
- Add CLI command for proposals [\#128](https://github.com/mysteriumnetwork/node/pull/128) ([donce](https://github.com/donce))
- Show tests success or failure message [\#127](https://github.com/mysteriumnetwork/node/pull/127) ([donce](https://github.com/donce))
- Feature/myst-259 validate session and signature on service provider side [\#126](https://github.com/mysteriumnetwork/node/pull/126) ([tadovas](https://github.com/tadovas))
- MYST-272 CLI for connection statistics [\#124](https://github.com/mysteriumnetwork/node/pull/124) ([donce](https://github.com/donce))
- MYST-274 Session duration [\#123](https://github.com/mysteriumnetwork/node/pull/123) ([donce](https://github.com/donce))
- Add script for linting only modified files [\#122](https://github.com/mysteriumnetwork/node/pull/122) ([donce](https://github.com/donce))
- MYST-114 Extract Authentificator interface [\#121](https://github.com/mysteriumnetwork/node/pull/121) ([Waldz](https://github.com/Waldz))
- Add statistics endpoint with bytes sent/received [\#120](https://github.com/mysteriumnetwork/node/pull/120) ([donce](https://github.com/donce))
- MYST-268 Add endpoint and CLI command for current IP [\#118](https://github.com/mysteriumnetwork/node/pull/118) ([donce](https://github.com/donce))
- Fix country database file to be reachable in docker [\#117](https://github.com/mysteriumnetwork/node/pull/117) ([donce](https://github.com/donce))
- When creating identity, unlock it before registering [\#116](https://github.com/mysteriumnetwork/node/pull/116) ([donce](https://github.com/donce))
- Fix command run error handling, which sometimes results in infinite loop [\#115](https://github.com/mysteriumnetwork/node/pull/115) ([donce](https://github.com/donce))
- Fix runtime error when aborting program before http server started [\#114](https://github.com/mysteriumnetwork/node/pull/114) ([donce](https://github.com/donce))
- openvpn auth poc with some cli fixes [\#113](https://github.com/mysteriumnetwork/node/pull/113) ([zolia](https://github.com/zolia))
- Fix some of linter warnings [\#112](https://github.com/mysteriumnetwork/node/pull/112) ([donce](https://github.com/donce))
- Allow creating identity with passphrase using CLI [\#111](https://github.com/mysteriumnetwork/node/pull/111) ([donce](https://github.com/donce))
- Cors handling middleware improved to handle preflight cors requests [\#110](https://github.com/mysteriumnetwork/node/pull/110) ([tadovas](https://github.com/tadovas))
- Change "password" term with "passphrase" to unify naming [\#109](https://github.com/mysteriumnetwork/node/pull/109) ([donce](https://github.com/donce))
- Make --cli optional when running  bin/client\_run [\#108](https://github.com/mysteriumnetwork/node/pull/108) ([donce](https://github.com/donce))
- MYST-148 gracefully close on Interrupt, sigterm and sighup [\#107](https://github.com/mysteriumnetwork/node/pull/107) ([shroomist](https://github.com/shroomist))
- Makefile: add build docker image commands [\#106](https://github.com/mysteriumnetwork/node/pull/106) ([zyfdegh](https://github.com/zyfdegh))
- WIP: add Makefile to install glide, build server and client [\#105](https://github.com/mysteriumnetwork/node/pull/105) ([zyfdegh](https://github.com/zyfdegh))
- Fix linter errors in "command\_run.identity" and "location" [\#104](https://github.com/mysteriumnetwork/node/pull/104) ([donce](https://github.com/donce))
- Fix failing tests due to package renaming [\#103](https://github.com/mysteriumnetwork/node/pull/103) ([donce](https://github.com/donce))
- Rename identity handlers to fix linter warnings [\#102](https://github.com/mysteriumnetwork/node/pull/102) ([donce](https://github.com/donce))
- Add country detection, send country in proposal [\#101](https://github.com/mysteriumnetwork/node/pull/101) ([donce](https://github.com/donce))
- Identity passphrase \(Unlock endpoint for tequilapi, server command option, CLI\) [\#100](https://github.com/mysteriumnetwork/node/pull/100) ([donce](https://github.com/donce))
- Fix request error handling to avoid segmentation violation [\#99](https://github.com/mysteriumnetwork/node/pull/99) ([donce](https://github.com/donce))
- MYST-114 Fix linter rules in "communication" package [\#97](https://github.com/mysteriumnetwork/node/pull/97) ([Waldz](https://github.com/Waldz))
- hex encoded signature changed to base64 [\#96](https://github.com/mysteriumnetwork/node/pull/96) ([zolia](https://github.com/zolia))
- Add cors headers for tequilapi service [\#95](https://github.com/mysteriumnetwork/node/pull/95) ([tadovas](https://github.com/tadovas))
- Refactor fake identity manager, fix test [\#94](https://github.com/mysteriumnetwork/node/pull/94) ([donce](https://github.com/donce))
- MYST-114 sign all mysterium api rest calls [\#93](https://github.com/mysteriumnetwork/node/pull/93) ([tadovas](https://github.com/tadovas))
- MYST-188 / CLI Client [\#92](https://github.com/mysteriumnetwork/node/pull/92) ([ignasbernotas](https://github.com/ignasbernotas))
- feature/MYST-114 sign register identity request [\#91](https://github.com/mysteriumnetwork/node/pull/91) ([tadovas](https://github.com/tadovas))
- MYST-228 Permit empty identity password when creating identity [\#90](https://github.com/mysteriumnetwork/node/pull/90) ([donce](https://github.com/donce))
- MYST-224 Refactor `identityHandler`, add tests [\#89](https://github.com/mysteriumnetwork/node/pull/89) ([donce](https://github.com/donce))
- Feature/myst-114-sign-mysterium-apimessages [\#88](https://github.com/mysteriumnetwork/node/pull/88) ([tadovas](https://github.com/tadovas))
- MYST-114 CommunicationChannel. Sign all payload messages [\#87](https://github.com/mysteriumnetwork/node/pull/87) ([Waldz](https://github.com/Waldz))
- Feature/myst 27 broker on public domain [\#86](https://github.com/mysteriumnetwork/node/pull/86) ([zolia](https://github.com/zolia))
- MYST-163 Add `golint` package [\#84](https://github.com/mysteriumnetwork/node/pull/84) ([donce](https://github.com/donce))
- Report process id of client in healthcheck endpoint [\#83](https://github.com/mysteriumnetwork/node/pull/83) ([tadovas](https://github.com/tadovas))
- HOTFIX Drop SessionStatsDeprecated DTO [\#82](https://github.com/mysteriumnetwork/node/pull/82) ([Waldz](https://github.com/Waldz))
- HOTFIX Stop growing NAT forwarding rules [\#81](https://github.com/mysteriumnetwork/node/pull/81) ([Waldz](https://github.com/Waldz))
- Improve bytescount middleware tests [\#80](https://github.com/mysteriumnetwork/node/pull/80) ([donce](https://github.com/donce))
- Improvement - print port number when started serving api requests [\#79](https://github.com/mysteriumnetwork/node/pull/79) ([tadovas](https://github.com/tadovas))
- hotfix - identities.List\(\) to return object instead of array [\#78](https://github.com/mysteriumnetwork/node/pull/78) ([shroomist](https://github.com/shroomist))
- MYST-182 Discovery. Use new stats endpoint [\#77](https://github.com/mysteriumnetwork/node/pull/77) ([donce](https://github.com/donce))
- myst-177 proposal list [\#76](https://github.com/mysteriumnetwork/node/pull/76) ([shroomist](https://github.com/shroomist))
- HOTFIX Node creates identity on each run [\#75](https://github.com/mysteriumnetwork/node/pull/75) ([Waldz](https://github.com/Waldz))
- MYST-114 / Message signing and verification [\#74](https://github.com/mysteriumnetwork/node/pull/74) ([ignasbernotas](https://github.com/ignasbernotas))
- Fixed issues with race conditions [\#73](https://github.com/mysteriumnetwork/node/pull/73) ([tadovas](https://github.com/tadovas))
- MYST-179 Create Docker container with minimal Alpine image [\#72](https://github.com/mysteriumnetwork/node/pull/72) ([Waldz](https://github.com/Waldz))
- MYST-143 tequila disconnect [\#71](https://github.com/mysteriumnetwork/node/pull/71) ([shroomist](https://github.com/shroomist))
- myst-161 connection status connecting [\#70](https://github.com/mysteriumnetwork/node/pull/70) ([shroomist](https://github.com/shroomist))
- MYST-181 Use proposals endpoint instead of create-session [\#69](https://github.com/mysteriumnetwork/node/pull/69) ([donce](https://github.com/donce))
- Fix connection creation failure handling [\#68](https://github.com/mysteriumnetwork/node/pull/68) ([donce](https://github.com/donce))
- Refactored tequilapi server to bind on port on start serving instead … [\#67](https://github.com/mysteriumnetwork/node/pull/67) ([tadovas](https://github.com/tadovas))
- MYST-173 securing openvpn [\#66](https://github.com/mysteriumnetwork/node/pull/66) ([zolia](https://github.com/zolia))
- MYST-102 client connection manager [\#65](https://github.com/mysteriumnetwork/node/pull/65) ([tadovas](https://github.com/tadovas))
- MYST-20 Mock NATS connection it tests [\#64](https://github.com/mysteriumnetwork/node/pull/64) ([Waldz](https://github.com/Waldz))
- removed unused imports [\#63](https://github.com/mysteriumnetwork/node/pull/63) ([ignasbernotas](https://github.com/ignasbernotas))
- myst-100 tequilapi create identity / register identity [\#62](https://github.com/mysteriumnetwork/node/pull/62) ([shroomist](https://github.com/shroomist))
- Add configuration step to CONTRIBUTING.md [\#61](https://github.com/mysteriumnetwork/node/pull/61) ([donce](https://github.com/donce))
- MYST-153 / Refactored dto.Identity string to identity.Identity struct [\#60](https://github.com/mysteriumnetwork/node/pull/60) ([ignasbernotas](https://github.com/ignasbernotas))
- fixed identity cache tests [\#59](https://github.com/mysteriumnetwork/node/pull/59) ([ignasbernotas](https://github.com/ignasbernotas))
- MYST-102 client api connect validation [\#58](https://github.com/mysteriumnetwork/node/pull/58) ([tadovas](https://github.com/tadovas))
- Bind local api on localhost interface by default - slight security im… [\#57](https://github.com/mysteriumnetwork/node/pull/57) ([tadovas](https://github.com/tadovas))
- fixed tests asking for network permission [\#56](https://github.com/mysteriumnetwork/node/pull/56) ([ignasbernotas](https://github.com/ignasbernotas))
- MYST-101 client. list identities teqilapi [\#55](https://github.com/mysteriumnetwork/node/pull/55) ([shroomist](https://github.com/shroomist))
- MYST-159 Node. Call API method on creating new identity [\#54](https://github.com/mysteriumnetwork/node/pull/54) ([ignasbernotas](https://github.com/ignasbernotas))
- MYST-81 reuse node identity [\#53](https://github.com/mysteriumnetwork/node/pull/53) ([ignasbernotas](https://github.com/ignasbernotas))
- MYST-99 bootstrap client api [\#52](https://github.com/mysteriumnetwork/node/pull/52) ([tadovas](https://github.com/tadovas))
- HOTFIX. Commit helper aliases [\#49](https://github.com/mysteriumnetwork/node/pull/49) ([Waldz](https://github.com/Waldz))
- MYST-80 Finished IdentityManager class \(generate identity, list identities\) [\#48](https://github.com/mysteriumnetwork/node/pull/48) ([interro](https://github.com/interro))
- MYST-41 Session generation and exchange [\#47](https://github.com/mysteriumnetwork/node/pull/47) ([ignasbernotas](https://github.com/ignasbernotas))
- MYST-54 Communication Channel. Create VPN config during session start  [\#46](https://github.com/mysteriumnetwork/node/pull/46) ([Waldz](https://github.com/Waldz))
- VPN client connection string transportation via NATS [\#45](https://github.com/mysteriumnetwork/node/pull/45) ([ignasbernotas](https://github.com/ignasbernotas))
- all: gofmt -w s [\#44](https://github.com/mysteriumnetwork/node/pull/44) ([fjl](https://github.com/fjl))
- Added config variable replacement from .env file during builds [\#43](https://github.com/mysteriumnetwork/node/pull/43) ([ignasbernotas](https://github.com/ignasbernotas))
- MYST-20 Communication dialog request/Response DTO models [\#40](https://github.com/mysteriumnetwork/node/pull/40) ([Waldz](https://github.com/Waldz))
- MYST-88/MYST-89 Service proposal serialisation  [\#39](https://github.com/mysteriumnetwork/node/pull/39) ([ignasbernotas](https://github.com/ignasbernotas))
- Node registration sends service proposal [\#38](https://github.com/mysteriumnetwork/node/pull/38) ([ignasbernotas](https://github.com/ignasbernotas))
- fix client package name [\#37](https://github.com/mysteriumnetwork/node/pull/37) ([vv1133](https://github.com/vv1133))
- MYST-39 Deserialise ServiceProposal DTOs from JSON [\#36](https://github.com/mysteriumnetwork/node/pull/36) ([Waldz](https://github.com/Waldz))
- Improve contribution documentation [\#35](https://github.com/mysteriumnetwork/node/pull/35) ([Waldz](https://github.com/Waldz))
- MYST-40 Communication Channel. Client creates dialog with Node [\#34](https://github.com/mysteriumnetwork/node/pull/34) ([Waldz](https://github.com/Waldz))
- Scripts for IP info retrieval \(ASN, ASN path, subnets, peer info\) [\#33](https://github.com/mysteriumnetwork/node/pull/33) ([Waldz](https://github.com/Waldz))
- MYST-18 define client promise [\#32](https://github.com/mysteriumnetwork/node/pull/32) ([interro](https://github.com/interro))
- MYST-6 Define ServiceProposal data structures [\#31](https://github.com/mysteriumnetwork/node/pull/31) ([Waldz](https://github.com/Waldz))
- Custom Communication Channel, Glide integration [\#30](https://github.com/mysteriumnetwork/node/pull/30) ([Waldz](https://github.com/Waldz))
- Building system [\#29](https://github.com/mysteriumnetwork/node/pull/29) ([Waldz](https://github.com/Waldz))
- Node traffic monitoring command [\#25](https://github.com/mysteriumnetwork/node/pull/25) ([Waldz](https://github.com/Waldz))
- use filepath.Join to concat path; fix loop err [\#21](https://github.com/mysteriumnetwork/node/pull/21) ([zyfdegh](https://github.com/zyfdegh))
- Fix syntax error in Dockerfile \#9 [\#15](https://github.com/mysteriumnetwork/node/pull/15) ([Waldz](https://github.com/Waldz))
- node docker update [\#12](https://github.com/mysteriumnetwork/node/pull/12) ([Goodsmileduck](https://github.com/Goodsmileduck))
- Traffic statistics "Data transfered" for node sessions \#4 [\#6](https://github.com/mysteriumnetwork/node/pull/6) ([Waldz](https://github.com/Waldz))

## [0.0.6](https://github.com/mysteriumnetwork/node/tree/0.0.6) (2017-05-09)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.0.5...0.0.6)

## [0.0.5](https://github.com/mysteriumnetwork/node/tree/0.0.5) (2017-05-08)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.0.4...0.0.5)

## [0.0.4](https://github.com/mysteriumnetwork/node/tree/0.0.4) (2017-05-08)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.0.3...0.0.4)

## [0.0.3](https://github.com/mysteriumnetwork/node/tree/0.0.3) (2017-05-07)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.0.2...0.0.3)

## [0.0.2](https://github.com/mysteriumnetwork/node/tree/0.0.2) (2017-05-04)
[Full Changelog](https://github.com/mysteriumnetwork/node/compare/0.0.1...0.0.2)

**Merged pull requests:**

- API client for sending stats [\#2](https://github.com/mysteriumnetwork/node/pull/2) ([Waldz](https://github.com/Waldz))

## [0.0.1](https://github.com/mysteriumnetwork/node/tree/0.0.1) (2017-05-01)
**Merged pull requests:**

- Possibility to test without dedicated server for node [\#1](https://github.com/mysteriumnetwork/node/pull/1) ([Waldz](https://github.com/Waldz))



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*
