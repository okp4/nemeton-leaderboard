# ØKP4 🧙 Nemeton Leaderboard 🏆

## [2.7.1](https://github.com/okp4/nemeton-leaderboard/compare/v2.7.0...v2.7.1) (2023-08-30)


### Bug Fixes

* **validator:** remove comment mistake ([f9f6fca](https://github.com/okp4/nemeton-leaderboard/commit/f9f6fca9d02fdb9eddd1d8e2b4ee8b68e12ce5e4))

## [2.7.0](https://github.com/okp4/nemeton-leaderboard/compare/v2.6.0...v2.7.0) (2023-08-11)


### Features

* **upgrade:** allow multiple upgrade task ([6a99ce7](https://github.com/okp4/nemeton-leaderboard/commit/6a99ce728bad5f7ef5016bd0979b561a4c78a639))


### Bug Fixes

* **lint:** bump golang lint to 1.53 ([74a2642](https://github.com/okp4/nemeton-leaderboard/commit/74a2642c47cf87821fd98c9d32982e2f224069dd))
* **lint:** fix new linter issue ([030820f](https://github.com/okp4/nemeton-leaderboard/commit/030820f7cdad38b2153f444fc937c3bc5e52e05a))
* **lint:** wastedassign linter issue ([9d8a223](https://github.com/okp4/nemeton-leaderboard/commit/9d8a223b460370de0caa65d6ecaacbc8ab2d7e90))
* **validator:** allocate initial zero point on create new validator ([49a06dc](https://github.com/okp4/nemeton-leaderboard/commit/49a06dcc76c590e2f3a24c72cf50356d0881725b))

## [2.6.0](https://github.com/okp4/nemeton-leaderboard/compare/v2.5.1...v2.6.0) (2023-05-24)


### Features

* add phase 5 tasks ([35b0ed6](https://github.com/okp4/nemeton-leaderboard/commit/35b0ed6bc4b921da614fb2ee9f388d31f51c90cb))

## [2.5.1](https://github.com/okp4/nemeton-leaderboard/compare/v2.5.0...v2.5.1) (2023-04-24)


### Bug Fixes

* avoid crash when no phase in progress ([5619aa5](https://github.com/okp4/nemeton-leaderboard/commit/5619aa5392fb7cdfb391ee1d41245e2e7d686667))
* **phase:** simplify end phase handling ([b0237bb](https://github.com/okp4/nemeton-leaderboard/commit/b0237bbe3f10a06f9d81ac3674a33dfacec7ac82))

## [2.5.0](https://github.com/okp4/nemeton-leaderboard/compare/v2.4.0...v2.5.0) (2023-02-28)


### Features

* **nemeton:** add phase 4 bootstrap ([0f4161e](https://github.com/okp4/nemeton-leaderboard/commit/0f4161ecec8107865696c463d6bcd047c83bec04))


### Bug Fixes

* **governance:** check also if validator vote with old MsgVote proto ([863126d](https://github.com/okp4/nemeton-leaderboard/commit/863126d6702bb0d1333b6f489b55382ded7768e2))
* **nemeton:** correct phase 3 bootstrap ([b5b7a59](https://github.com/okp4/nemeton-leaderboard/commit/b5b7a59bece3d58560537a458609b1d54f613773))

## [2.4.0](https://github.com/okp4/nemeton-leaderboard/compare/v2.3.0...v2.4.0) (2023-02-20)


### Features

* allow override completed task ([75ac78c](https://github.com/okp4/nemeton-leaderboard/commit/75ac78cda52be16d9b2a6e777488c5bd9478554a))
* **bonus:** add event for bonus points submitted ([7600b32](https://github.com/okp4/nemeton-leaderboard/commit/7600b321872b521d2217fc30b144d9213bd85873))
* **bonus:** add graphql mutation for submit bonus points ([248cfd1](https://github.com/okp4/nemeton-leaderboard/commit/248cfd18334f6e2dd6f3461f196abe829fc68112))
* **bonus:** include bonus point into validator graphql api ([a17e44a](https://github.com/okp4/nemeton-leaderboard/commit/a17e44adf5f6062068c8690357885b1373e0001c))
* **phase3:** details the tasks ([ffb6274](https://github.com/okp4/nemeton-leaderboard/commit/ffb6274520b7978bd079a8ed522431ce2f1a5bb4))
* **upgrade:** add a new upgrade task check ([d0607a9](https://github.com/okp4/nemeton-leaderboard/commit/d0607a97700d210a4405f7e84e329f8f3dab1f80))
* **upgrade:** complete upgrade task if validator sign after height given ([faf5cb0](https://github.com/okp4/nemeton-leaderboard/commit/faf5cb0bda2c251409869c605b5d1ce8a38d8d8e))


### Bug Fixes

* **bonus:** make bonusPoints attributes mandatory on api schema ([77ae2bc](https://github.com/okp4/nemeton-leaderboard/commit/77ae2bcca795f5a8059953b4f2cf2c121a262774))
* **subscription:** add return to make linter happy ([22e9fb3](https://github.com/okp4/nemeton-leaderboard/commit/22e9fb36ae6d10cd385600bb923c968e1dc2667b))
* **subscription:** panic if no phase was found for actual block range ([5f38875](https://github.com/okp4/nemeton-leaderboard/commit/5f38875457bf405f0c56f392ec9438283617ade7))

## [2.3.0](https://github.com/okp4/nemeton-leaderboard/compare/v2.2.0...v2.3.0) (2023-02-01)


### Features

* **sync:** register MsgVote messages ([10660c9](https://github.com/okp4/nemeton-leaderboard/commit/10660c97ddae07e39d33a03d68fbec5cebacb930))
* **vote:** create vote proposal task type ([6f26190](https://github.com/okp4/nemeton-leaderboard/commit/6f26190914639b82368b7457079496826cafb16d))
* **vote:** get param proposal helper ([73a0524](https://github.com/okp4/nemeton-leaderboard/commit/73a0524cd259a80edb05c944dde1072a5c6531f7))


### Bug Fixes

* **lint:** fix lint issue ([e8a7d93](https://github.com/okp4/nemeton-leaderboard/commit/e8a7d938962f5dead5ca7cadb939e126c5d67d94))

## [2.2.0](https://github.com/okp4/nemeton-leaderboard/compare/v2.1.1...v2.2.0) (2023-01-11)


### Features

* **task:** add submitted status to tasks with submission ([edd0e2b](https://github.com/okp4/nemeton-leaderboard/commit/edd0e2be582375ba911a1b2ccada2bae424e9135))
* **task:** allow submitting a task through mutation ([4afe843](https://github.com/okp4/nemeton-leaderboard/commit/4afe8436140b427d23490ebaebc1d441b76a2251))
* **validator:** add mutation to remove validator ([1e92a2e](https://github.com/okp4/nemeton-leaderboard/commit/1e92a2e821148caffd5ee0710a17911102411a6a))
* **validator:** handle remove validator event ([237740b](https://github.com/okp4/nemeton-leaderboard/commit/237740b84d98bbffe7c862f6b487d317f620d06d))
* **validator:** trigger ValidatorRemovedEvent ([fbfb01a](https://github.com/okp4/nemeton-leaderboard/commit/fbfb01ad90a46415c79c7065a06e74f1c547cc84))


### Bug Fixes

* **task:** make manual task completion intemporal ([f60bc68](https://github.com/okp4/nemeton-leaderboard/commit/f60bc6884fa2631b6f4bc0d318b90464f89748c4))

## [2.1.1](https://github.com/okp4/nemeton-leaderboard/compare/v2.1.0...v2.1.1) (2022-12-23)


### Bug Fixes

* **graphql:** repair points marshalling ([dcf3a10](https://github.com/okp4/nemeton-leaderboard/commit/dcf3a108c3fa27a3e2d588b0c44574a3a02e7101))

## [2.1.0](https://github.com/okp4/nemeton-leaderboard/compare/v2.0.0...v2.1.0) (2022-12-23)


### Features

* **phase2:** add dashboard and snapshot field on validator ([ee8cffc](https://github.com/okp4/nemeton-leaderboard/commit/ee8cffcf6a49903f2ddd8435f83db5bb65582600))
* **phase2:** add maxpoints on dashboard task ([f68baa9](https://github.com/okp4/nemeton-leaderboard/commit/f68baa9e9adfae2842fda854ee627edf7c6f400d))
* **phase2:** allow add rewards on dashboard mutation ([b4a9072](https://github.com/okp4/nemeton-leaderboard/commit/b4a9072fb1609eab01793c2eb13c7ec7683b268c))
* **phase2:** register dashboard rewards onmongo ([d45b1df](https://github.com/okp4/nemeton-leaderboard/commit/d45b1df3001d81b104fbb76f24969cc1cbf225f2))


### Bug Fixes

* **event:** add unsubscribe event to stop mongo stream on crash ([daf5fa1](https://github.com/okp4/nemeton-leaderboard/commit/daf5fa1ae7f35c3c2030b5d8c2e2991878547159))
* **event:** remove actor pid stream on stop ([bd2942b](https://github.com/okp4/nemeton-leaderboard/commit/bd2942b5bbea9ad861da0d71ac463e171e1ceb84))
* **lint:** fix linter and naming ([54905aa](https://github.com/okp4/nemeton-leaderboard/commit/54905aa1bc4ce1a5a16427c60d6ef53a65e01cc9))
* **register:** ensure missedblocks propulated ([2e3dc96](https://github.com/okp4/nemeton-leaderboard/commit/2e3dc96133f15e80cdac6c9e6cbc7c863c3c5f14))
* **update:** prevent erasing tasks & points on update ([148554c](https://github.com/okp4/nemeton-leaderboard/commit/148554c5a514fdce46e5bbb02f59038c03ec7f11))

## [2.0.0](https://github.com/okp4/nemeton-leaderboard/compare/v1.3.0...v2.0.0) (2022-12-22)


### ⚠ BREAKING CHANGES

* **graphql:** simplify schema

### Features

* **event:** design task completed event ([11a836b](https://github.com/okp4/nemeton-leaderboard/commit/11a836b522309697ab76a817794f140ceb0e00fd))
* **event:** design validator updated event ([a10a34c](https://github.com/okp4/nemeton-leaderboard/commit/a10a34c659bbfb5c3697c0d2e24ec619c723c68a))
* **event:** implements validator registered evt publish ([6da11ed](https://github.com/okp4/nemeton-leaderboard/commit/6da11ed47a2404ada38ce942fb581e402934708f))
* **graphql:** implement completeTask mutation ([ce2faf3](https://github.com/okp4/nemeton-leaderboard/commit/ce2faf31c4f3a9cc4a1c66df5597ef775cf4c17a))
* **graphql:** implement updateValidator mutation ([1176b71](https://github.com/okp4/nemeton-leaderboard/commit/1176b715c9c9b01f866ced6841804a72bc1a2c0e))
* **graphql:** introduce validator registration mutation ([08d09b6](https://github.com/okp4/nemeton-leaderboard/commit/08d09b6e793589ff0860df12fdb581f68ba0d67e))
* **graphql:** simplify schema ([0893f89](https://github.com/okp4/nemeton-leaderboard/commit/0893f89d96063093129f5aeefc30a295fb9abde9))
* **graphql:** specify completeTask mutation ([48f54a7](https://github.com/okp4/nemeton-leaderboard/commit/48f54a7dc08daed24330d5939dc4fc017af403db))
* **graphql:** specify updateValidator mutation ([eb08434](https://github.com/okp4/nemeton-leaderboard/commit/eb08434e4b4748ed9a4e882a7d279af67fc74c08))
* **grpc:** implements validator fetching ([f0d6b23](https://github.com/okp4/nemeton-leaderboard/commit/f0d6b237c8432cc4f185bae65cb2db07d406163f))
* **grpc:** share grpc client actor with graphql ([a9882e3](https://github.com/okp4/nemeton-leaderboard/commit/a9882e3b4aa901bf31b72683ea07f007068feadf))
* **leaderboard:** handle task completed event ([e28027d](https://github.com/okp4/nemeton-leaderboard/commit/e28027d97961a5df06464b6fe450cb46c11aca51))
* **leaderboard:** handle validator updated event ([c86e862](https://github.com/okp4/nemeton-leaderboard/commit/c86e862fdd3880c7f0b0397cee1924b9c829f0ea))
* **leaderboard:** implements validator register ([a92e3b9](https://github.com/okp4/nemeton-leaderboard/commit/a92e3b90f8695b0778934e243b6f7066c4101cf4))
* **nemeton:** bootstrap phase2 ([1d34833](https://github.com/okp4/nemeton-leaderboard/commit/1d3483325bc76a4d38fdff2c0eb617b665e1814f))
* **nemeton:** introduce new task types ([59d41eb](https://github.com/okp4/nemeton-leaderboard/commit/59d41eba231103f7665ad249dae133b1625758b1))
* **phase2:** add mutation to register rpc ([340e1b0](https://github.com/okp4/nemeton-leaderboard/commit/340e1b069db5c7beb545bebb7a551e708285b271))
* **phase2:** handle Register rpc event ([2a4d0af](https://github.com/okp4/nemeton-leaderboard/commit/2a4d0af236f11f1cfc3c5b4b7bba05250333c322))
* **phase2:** include rpc endpoint in mongo validator model ([702d9e8](https://github.com/okp4/nemeton-leaderboard/commit/702d9e878d65ee85a0880407b7bcdd4f4c7faa53))
* **phase2:** omit empty rpc endpoint field ([f6c4730](https://github.com/okp4/nemeton-leaderboard/commit/f6c473008279e34031f654e08ab1fc6d04a4ef12))
* **phase2:** register event for rpc endpoint ([4ce3fa1](https://github.com/okp4/nemeton-leaderboard/commit/4ce3fa17eeaa45ef6338774a17f1ff81df81ce42))
* **phase2:** register reward for rpc endpoint ([555a177](https://github.com/okp4/nemeton-leaderboard/commit/555a177b3e36329142349aa963a833fa3efc6100))
* **phase:** handle phase started and phase end ([7b3013b](https://github.com/okp4/nemeton-leaderboard/commit/7b3013bd05f3492257db4876afe0d9b5e016652e))
* **phase:** register validator uptime points on phase ended ([1b2f16f](https://github.com/okp4/nemeton-leaderboard/commit/1b2f16fc02474cfe35dab2bccb934dd78e9d22b3))
* **uptime:** get uptime max point from task ([85d6252](https://github.com/okp4/nemeton-leaderboard/commit/85d62521601da23135d77cb2f1394e653907f2ef))


### Bug Fixes

* **lint:** linter correction ([c711a9f](https://github.com/okp4/nemeton-leaderboard/commit/c711a9fb692fb066e1baa10e35dab8235f677a21))
* **log:** talk about the right thing ([9f6d06e](https://github.com/okp4/nemeton-leaderboard/commit/9f6d06eff4f3bbc9604fa525eb1b2968e1c6d865))
* **phase2:** register rpc by valoper instead of moniker ([9ecfd63](https://github.com/okp4/nemeton-leaderboard/commit/9ecfd637077fb760bebe871ee1e99922d82d6bae))
* **phase2:** remove old graphql generated code ([5de7dc0](https://github.com/okp4/nemeton-leaderboard/commit/5de7dc0fa44b2e1c0cbbe1e974a8eed5a3595921))
* **phase:** get correct previous phase ([d33b092](https://github.com/okp4/nemeton-leaderboard/commit/d33b0924625b8e4bd2c285c186cf7a8bc0448364))
* **phase:** uptime point round ([480a887](https://github.com/okp4/nemeton-leaderboard/commit/480a887271d67bdfd214f8df67d5bb2ba5b28d32))
* **sync:** correct inconcistant log message ([41149fc](https://github.com/okp4/nemeton-leaderboard/commit/41149fc859e6486b8e4046a671bbffaf84cbdc52))
* **task:** cast uptime from int64 to uint64 ([993a891](https://github.com/okp4/nemeton-leaderboard/commit/993a891398a53f9c6131d90f551ce7dfc304d994))

## [1.3.0](https://github.com/okp4/nemeton-leaderboard/compare/v1.2.0...v1.3.0) (2022-12-15)


### Features

* **graphql:** return all phases in per phase validator tasks ([d3762c8](https://github.com/okp4/nemeton-leaderboard/commit/d3762c8215c70c588a3efeb4f241200ad1ba14bc))
* **graphql:** return the points earned per phase ([86e650a](https://github.com/okp4/nemeton-leaderboard/commit/86e650ae15439ab908986e9d921d1ee528f910d9))
* **logging:** log essential infos on new block event handling ([52af41b](https://github.com/okp4/nemeton-leaderboard/commit/52af41b87e52119b45020194a7f2a98813af4b10))
* **task:** implements node setup task completion ([7ca6a1d](https://github.com/okp4/nemeton-leaderboard/commit/7ca6a1dc65bc68156b7dc1a0a1002b9cccae409f))

## [1.2.0](https://github.com/okp4/nemeton-leaderboard/compare/v1.1.1...v1.2.0) (2022-12-14)


### Features

* **phase:** add query to update phase block range ([27da464](https://github.com/okp4/nemeton-leaderboard/commit/27da4643009022181e67bf100ff05e8a6c973c2e))
* **phase:** return block range on graphql phase ([35260fd](https://github.com/okp4/nemeton-leaderboard/commit/35260fdc2aee56a7c59fcf058bc3c182ead5dead))
* **phase:** update phase block range on new block ([99880d9](https://github.com/okp4/nemeton-leaderboard/commit/99880d994ea1fd2f52524db9caf356dad1ba6ccb))
* **tweet:** handle tweet event ([1f083a0](https://github.com/okp4/nemeton-leaderboard/commit/1f083a0baf935940dc43f8021e9f0ab8b99dac57))
* **tweet:** increment point if task is validated ([e37c48e](https://github.com/okp4/nemeton-leaderboard/commit/e37c48ee9719bdb0eec8b5935525732d5c23cf47))


### Bug Fixes

* **ci:** fix lint go ([3c9a0e0](https://github.com/okp4/nemeton-leaderboard/commit/3c9a0e0e4e7d56824d2513adabda55b74f70da73))
* **ci:** make linter happy ([9d26039](https://github.com/okp4/nemeton-leaderboard/commit/9d26039659e4c9bbcc075914123f4fe8f2368f1a))
* **graphql:** dissociate phase blocks fetch from phase ([7047db9](https://github.com/okp4/nemeton-leaderboard/commit/7047db9f67149e687f847d5c7291ff82353c1fa1))
* **graphql:** make phase block range optional ([8d96bb2](https://github.com/okp4/nemeton-leaderboard/commit/8d96bb2cf8ba7a63cd1af1e25be4247d6509ddbc))
* **phase:** incorrect date for phase 2 ([370eb81](https://github.com/okp4/nemeton-leaderboard/commit/370eb814734f1d365d547d49dc2fb9c96e34bea3))
* **phase:** make block range to exclusif ([6aec212](https://github.com/okp4/nemeton-leaderboard/commit/6aec2127f83c4833edce4f2c17293a5bb45e93c4))
* **store:** handle event in their time context ([4bf6879](https://github.com/okp4/nemeton-leaderboard/commit/4bf6879c706ff0299c6d1d411bd03a5612a820c3))
* **tasks:** update first task end date ([ed1be39](https://github.com/okp4/nemeton-leaderboard/commit/ed1be3914c92f2436814a974eb1e1ecb1c427b48))
* **tweet:** handle tweet on event date instead of now ([b90efdc](https://github.com/okp4/nemeton-leaderboard/commit/b90efdc58fc533c891b025a88aa463bb6210e096))
* **tweet:** only earn point if tweet has been posted on task period ([84d2426](https://github.com/okp4/nemeton-leaderboard/commit/84d2426a8c5672330570304ed3f629ad5df6d642))

## [1.1.1](https://github.com/okp4/nemeton-leaderboard/compare/v1.1.0...v1.1.1) (2022-12-06)


### Bug Fixes

* **mongo:** avoid client connection leak ([692be49](https://github.com/okp4/nemeton-leaderboard/commit/692be49cb7459cb7b15ba57396228068275ee668))

## [1.1.0](https://github.com/okp4/nemeton-leaderboard/compare/v1.0.1...v1.1.0) (2022-12-05)


### Features

* **auth:** add middleware to extract bearer in ctx ([1271293](https://github.com/okp4/nemeton-leaderboard/commit/127129355d0423a1d48bc2990b3361790e1830b1))
* **auth:** introduce [@auth](https://github.com/auth) directive on submit gentx ([dec14b9](https://github.com/okp4/nemeton-leaderboard/commit/dec14b97a68399fd3101487a4f60d9c3159c7f63))
* **gentx:** implements gentx parsing to extract val ([e6e983d](https://github.com/okp4/nemeton-leaderboard/commit/e6e983d28f1a6af3ad252326b7946a246388e7f0))
* **gentx:** implements gentx submit mutation evt sending ([889e215](https://github.com/okp4/nemeton-leaderboard/commit/889e215f2974b63ac99fe9b147be5d459a2dac7e))
* **gentx:** implements gentx submitted event handling ([40f4d8e](https://github.com/okp4/nemeton-leaderboard/commit/40f4d8e9d5af6b0df1835cc65804c575816d4767))
* **graphql:** add mutation for gentx submission ([7c43217](https://github.com/okp4/nemeton-leaderboard/commit/7c43217d98e689ae970783255c5a33b73b371911))
* **graphql:** add started task count on validator ([ded2156](https://github.com/okp4/nemeton-leaderboard/commit/ded215660a6265392df0fb64e3d159a26d3bf8cc))
* **graphql:** add validator details fields ([65c97e0](https://github.com/okp4/nemeton-leaderboard/commit/65c97e0f2270bba483504921dec6d4b75c922bcf))
* **graphql:** implements json scalar ([9fc03d1](https://github.com/okp4/nemeton-leaderboard/commit/9fc03d1729dce506ebb2c2e5fcb405dcf55e2303))
* **mongo:** add bson codec for cosmos addr ([7cec838](https://github.com/okp4/nemeton-leaderboard/commit/7cec83831d978faa7eab86f7b6a1e19d8b881c63))
* **store:** add valcons index ([66939e9](https://github.com/okp4/nemeton-leaderboard/commit/66939e9ed4a2028289b1263b6a0fe1daf563792b))
* **store:** remove delegator filter in search ([2701c1e](https://github.com/okp4/nemeton-leaderboard/commit/2701c1e6e2d330a4307c38c27606074208de95c1))
* **subscription:** handle only NewBlockEvent ([fb5a0f0](https://github.com/okp4/nemeton-leaderboard/commit/fb5a0f0ad2ac3a5963f8211a33a589ed790e9aef))
* **subscription:** remove unnecessary block range struct from validator ([d284b53](https://github.com/okp4/nemeton-leaderboard/commit/d284b53bf6c031cf3c72cb8a3d1bae5be6cdee5a))
* **subscription:** save current event handler state ([924c1e2](https://github.com/okp4/nemeton-leaderboard/commit/924c1e2a74444f9e1c3588eb9259e14d0fdbaed9))
* **subscription:** subscribe to event new block ([7e6818c](https://github.com/okp4/nemeton-leaderboard/commit/7e6818cbb2088e93c312634ef1aa2d4a44856ff3))
* **subscription:** update uptime on nemeton store ([0266a9b](https://github.com/okp4/nemeton-leaderboard/commit/0266a9b2c8baeffd6894e72462c4dade5a7d1dd7))
* **tasks:** add gentx task type ([8871620](https://github.com/okp4/nemeton-leaderboard/commit/8871620d4e6d4b451274bc661122cc06f8245ad6))


### Bug Fixes

* **gentx:** handle nullity on msg fields ([323dbd7](https://github.com/okp4/nemeton-leaderboard/commit/323dbd7f97044a1bd3a06cfbf403b3819a09b538))
* **gentx:** remove website from mutation inputs ([ed0d254](https://github.com/okp4/nemeton-leaderboard/commit/ed0d254a8488fbde7967b57657a8cc7c8537ce67))
* **store:** avoid always fetching validators missedBlocks ([e5053e9](https://github.com/okp4/nemeton-leaderboard/commit/e5053e95ae2c9942ab04e1b7e0e6f68b526edc80))
* **subscription:** fix quering valoper address ([1d7d889](https://github.com/okp4/nemeton-leaderboard/commit/1d7d8891aadb5938b54761c9ccfde247904a5a08))
* **subscription:** match validator by cons addr instead of valoper ([5353e70](https://github.com/okp4/nemeton-leaderboard/commit/5353e703f296d8cc04f35570595d1fe895ae0314))
* **tweet:** make linter happy ([5252b98](https://github.com/okp4/nemeton-leaderboard/commit/5252b980a5a113c3c4f53cd60014a27d446909c2))
* **tweet:** use twitter account instead of hashtag for search tweet ([54928b2](https://github.com/okp4/nemeton-leaderboard/commit/54928b2088b1cb47a60ce531d1cbfa364c434a58))

## [1.0.1](https://github.com/okp4/nemeton-leaderboard/compare/v1.0.0...v1.0.1) (2022-11-30)


### Bug Fixes

* **sidh:** delay launch at 11am ([e3dbfc0](https://github.com/okp4/nemeton-leaderboard/commit/e3dbfc01b39f14cad5726679c05e213332b6847c))

## 1.0.0 (2022-11-30)


### Features

* add actor system init bootstrap ([44be149](https://github.com/okp4/nemeton-leaderboard/commit/44be149e5779ed635c2cc2788e1681dcfd94684e))
* add mongo indexes ([dcfbe0a](https://github.com/okp4/nemeton-leaderboard/commit/dcfbe0a6fc17229b6b9a2cd889f6ac7d2a071647))
* **board:** implement filter by search on validators ([db8841d](https://github.com/okp4/nemeton-leaderboard/commit/db8841dcafea511803d87bb2b8214ee661e10405))
* **cmd:** implements start cmd ([498acc1](https://github.com/okp4/nemeton-leaderboard/commit/498acc1ebf7d3565d3f0653c421ea45f725b9514))
* configure first phase bootstraping ([fc62b62](https://github.com/okp4/nemeton-leaderboard/commit/fc62b62727aff31f741fb8511f95c8933d03f9a1))
* **event:** add empty event store actor ([cf5fa70](https://github.com/okp4/nemeton-leaderboard/commit/cf5fa7064d306620f70f74feb6c0fd8aebb92e41))
* **event:** add event model ([1bcbc58](https://github.com/okp4/nemeton-leaderboard/commit/1bcbc58216b71ccc27d11befab65d67408f10f02))
* **event:** add store in actor ([c317273](https://github.com/okp4/nemeton-leaderboard/commit/c3172738cb406dfafdb039baa11bb1308a47f6f0))
* **event:** handle subscribe event message ([6698596](https://github.com/okp4/nemeton-leaderboard/commit/66985969cfc776cc86e85b353e050d54a8ed1934))
* **event:** implement event store stream ([2c57d8c](https://github.com/okp4/nemeton-leaderboard/commit/2c57d8cd966e3a3c49999a01b0a8e38205a5d55b))
* **event:** implement publish event ([1e5aadd](https://github.com/okp4/nemeton-leaderboard/commit/1e5aadd55441f09de681724a2f3787bd511731f2))
* **event:** implement stream actor ([8df900c](https://github.com/okp4/nemeton-leaderboard/commit/8df900c45863ad857963f85efd54fa751168498e))
* **event:** implements event store ([492ce4d](https://github.com/okp4/nemeton-leaderboard/commit/492ce4d77699e394a0c47e954c621976bd731282))
* **graphql:** add first schema ([0b465c7](https://github.com/okp4/nemeton-leaderboard/commit/0b465c7ec928a25d126862accecd72d62d306998))
* **graphql:** add test data ([90c3697](https://github.com/okp4/nemeton-leaderboard/commit/90c3697ecf88b893cf23c96a24802d3d3d1620c6))
* **graphql:** add type to tasks ([39023e9](https://github.com/okp4/nemeton-leaderboard/commit/39023e974e0a45e6fd20fe3c503248ff16dbc770))
* **graphql:** auto start server ([e99ed69](https://github.com/okp4/nemeton-leaderboard/commit/e99ed692e83f11fc78aa98bb7d750a337cac417d))
* **graphql:** document schema ([779389a](https://github.com/okp4/nemeton-leaderboard/commit/779389a2fa1c4119b601a657a222818c22296cc4))
* **graphql:** impl phases queries ([b7a9700](https://github.com/okp4/nemeton-leaderboard/commit/b7a970072ad82112ec150a24138f6483282ce2c2))
* **graphql:** implement get all phases ([8719ba7](https://github.com/okp4/nemeton-leaderboard/commit/8719ba77e14d9e743512525ca6a057f632afdec4))
* **graphql:** implement get phase query ([e7fc3e3](https://github.com/okp4/nemeton-leaderboard/commit/e7fc3e370de7f461be4090a05a81e34aa7e6df3d))
* **graphql:** implement validator count ([79e2a25](https://github.com/okp4/nemeton-leaderboard/commit/79e2a253af7451ba214952cb02080a7a9c651807))
* **graphql:** implement validator rank resolver ([082067f](https://github.com/okp4/nemeton-leaderboard/commit/082067fe60a7a0dddc1f15811969a73881ce0423))
* **graphql:** implements basic task state retrieval ([0939f70](https://github.com/okp4/nemeton-leaderboard/commit/0939f706d961496bf4e83cf58be2dde83c13c478))
* **graphql:** implements board connection ([1a4a7c7](https://github.com/okp4/nemeton-leaderboard/commit/1a4a7c7db7d847b3461f007f86bbe1c9f162a531))
* **graphql:** implements bsaci validator query ([88731cf](https://github.com/okp4/nemeton-leaderboard/commit/88731cf70d20b537bb846115b4b4ac9dba69e0f8))
* **graphql:** implements graphql actor ([83b9e7f](https://github.com/okp4/nemeton-leaderboard/commit/83b9e7f627c6801b90dd8ed5c13d7e9398e67c4b))
* **graphql:** implements http graphql server ([4f68445](https://github.com/okp4/nemeton-leaderboard/commit/4f6844559518bc8ea761515ecd21089187049a5d))
* **graphql:** implements KID scalar ([0ccbb02](https://github.com/okp4/nemeton-leaderboard/commit/0ccbb02c6eaf1abd976eaa03873c6044d44ffaca))
* **graphql:** implements scalars ([8f16c31](https://github.com/okp4/nemeton-leaderboard/commit/8f16c312bde9822ac3bab667df87ffab64c245ad))
* **graphql:** link resolver to custom phase model ([7fb46bf](https://github.com/okp4/nemeton-leaderboard/commit/7fb46bf0db079576a722983ee67f57df5cc0c09a))
* **graphql:** resolve temporal fields ([89b1939](https://github.com/okp4/nemeton-leaderboard/commit/89b19397db6f50e3bb1cd29249cc9e0be0706d80))
* **graphql:** specify int types and reduce default pagination ([f1e3340](https://github.com/okp4/nemeton-leaderboard/commit/f1e33402e971b7dad9ee1d9eca7926ce6b8d4459))
* **graphql:** switch to a base58 encoded cursor ([18e0b8b](https://github.com/okp4/nemeton-leaderboard/commit/18e0b8be9ab87eddff75da5c9ab3694c52f8601a))
* **graphql:** use a composite cursor ([a6646bc](https://github.com/okp4/nemeton-leaderboard/commit/a6646bc9e97cc5dee23c7ede27ac46fd4eb5e8cb))
* **graphql:** wire basic validator fields ([9c4a63f](https://github.com/okp4/nemeton-leaderboard/commit/9c4a63f12725100f9dddf44ec93736f18b9f8d3a))
* **graphql:** wire identity resolution with keybase client ([c5e59de](https://github.com/okp4/nemeton-leaderboard/commit/c5e59deff9c671eca3269d983c423bbab86e3da8))
* **graphql:** wire resolver with store ([7edec1f](https://github.com/okp4/nemeton-leaderboard/commit/7edec1f12031bfa318efbdc6f0c2f6cd78a010bb))
* **grpc:** add grpc actor client ([be23280](https://github.com/okp4/nemeton-leaderboard/commit/be23280854e802bdfbdca595ee614f221eb43c8b))
* **grpc:** add message to get the latest block ([65c5722](https://github.com/okp4/nemeton-leaderboard/commit/65c57227004435e02cfba24bafdafd6947160bf5))
* **grpc:** move all messages in the same file ([79c438b](https://github.com/okp4/nemeton-leaderboard/commit/79c438bedfd3e36d170d175cd059171d6fb24dde))
* introduce an offset storage ([2b4538f](https://github.com/okp4/nemeton-leaderboard/commit/2b4538f9e3ffeea709e11f9ff653ad924567f397))
* **keybase:** resolve identity picture ([2df169b](https://github.com/okp4/nemeton-leaderboard/commit/2df169b0f153e84e8b1a18555d9cf3ab538d8543))
* **phase:** implement phases bootstrap logic ([97d8c8c](https://github.com/okp4/nemeton-leaderboard/commit/97d8c8c64c95a3d5e52f9332c167e2b0b0706277))
* **phases:** add basic store and model for phases ([2f600d3](https://github.com/okp4/nemeton-leaderboard/commit/2f600d36fc9dfb5ae01f562e12d95f15cddc81d5))
* set fail strategies to system ([329f72a](https://github.com/okp4/nemeton-leaderboard/commit/329f72aa0b53e7cf2954d52a318fc973465eb63e))
* **sync:** add actor to sync block each 8 seconds ([655a0fd](https://github.com/okp4/nemeton-leaderboard/commit/655a0fdf52cf8c8f1bdac42dfb22bdbbc4c6d4c7))
* **sync:** cacth up block to latest block ([02b154e](https://github.com/okp4/nemeton-leaderboard/commit/02b154efc2da06b9c9a441d72567cacd3b04fb91))
* **sync:** publish event in store ([35753b6](https://github.com/okp4/nemeton-leaderboard/commit/35753b63a55d0ac26300d552f04ff5c393e263a1))
* **sync:** save sync state on store ([8506840](https://github.com/okp4/nemeton-leaderboard/commit/85068406c1693df8a1708c742d79fa84a0c778b6))
* **sync:** use actor scheduler instead of goroutine ([736f086](https://github.com/okp4/nemeton-leaderboard/commit/736f0864b27a8ffa516fdea289baa22e6fd55547))
* **tweet:** actor to fetch tweet at specified query with pagination ([64e2e6d](https://github.com/okp4/nemeton-leaderboard/commit/64e2e6d111cd8f154f981b63684ec801bd704a9b))
* **tweet:** publish new tweet event ([7143033](https://github.com/okp4/nemeton-leaderboard/commit/714303364adb07cf5dfe33a413c71022c62be8c7))
* **tweet:** save last tweet id in store ([fb52d7a](https://github.com/okp4/nemeton-leaderboard/commit/fb52d7a12d23bdb5f055095b67534661035aedc3))
* **tweet:** set the hashtag through command line args ([9ad11d1](https://github.com/okp4/nemeton-leaderboard/commit/9ad11d147e0c1fcb8c7e6f709cfac9bfa700da31))


### Bug Fixes

* avoid logging sensible informations ([f2573a0](https://github.com/okp4/nemeton-leaderboard/commit/f2573a0d972e15feb4314bb8184764b6c64534fb))
* **event:** make it work ([00ef051](https://github.com/okp4/nemeton-leaderboard/commit/00ef0513c102517cdb14b5aa0afa456b9792cf07))
* **graphql:** make validator tasks non nullable ([bb6ab27](https://github.com/okp4/nemeton-leaderboard/commit/bb6ab27d836333360c64ca4ea72bc2bf413a597c))
* **graphql:** properly marshal address scalars ([f422f6f](https://github.com/okp4/nemeton-leaderboard/commit/f422f6f8b1d9b9a268adce1502d9c22bd051be77))
* **lint:** make linter happy ([95d7857](https://github.com/okp4/nemeton-leaderboard/commit/95d78570a84ca0f56d54d2fdf94efc21584e30e5))
* **offset:** allow creating offset in not existing ([65c84f2](https://github.com/okp4/nemeton-leaderboard/commit/65c84f2fe6ed6f4f07c1dab5677648baf1e388cf))
* **sync:** add log and make linter happy ([f217ac9](https://github.com/okp4/nemeton-leaderboard/commit/f217ac99f22b947237b48e250143c681f9088786))
* **sync:** fix linter ([0f68dbb](https://github.com/okp4/nemeton-leaderboard/commit/0f68dbb3aa162742a80d0e02f31050ae784fe6bf))
* **tweet:** rename tweeter by twitter ([c349ec0](https://github.com/okp4/nemeton-leaderboard/commit/c349ec09c6d15412abd1f3d95675540b6965da56))
