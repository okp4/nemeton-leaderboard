# ØKP4 🧙 Nemeton Leaderboard 🏆

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
