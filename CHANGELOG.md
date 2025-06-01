# Changelog

## [1.1.0](https://github.com/meysam81/tarzan/compare/v1.0.2...v1.1.0) (2025-06-01)


### Features

* bring back the markdown support, no image though! ([f0bd113](https://github.com/meysam81/tarzan/commit/f0bd11377a2f832e467e1e8ebbeecf29f479249b))
* **build:** add hook for updating env example ([01986bf](https://github.com/meysam81/tarzan/commit/01986bf745a0cc2c91e517ed9095e70bc2bae9ff))
* **build:** create a production ready vite ([#5](https://github.com/meysam81/tarzan/issues/5)) ([65fa261](https://github.com/meysam81/tarzan/commit/65fa2619c712858fa0654b08ee40532641c99780))
* **build:** remove hardcoded URL to make the docker image portable ([c8bee63](https://github.com/meysam81/tarzan/commit/c8bee633409046fe64e2a3702fbed1607b833292))
* **build:** update env example on every commit ([7ebfacd](https://github.com/meysam81/tarzan/commit/7ebfacd8f2ac6029aeb9e9a2f8414a81dbe375f8))
* create 404 not found page ([cfe968c](https://github.com/meysam81/tarzan/commit/cfe968c5d502c5734e9fb49a777d8d8a8d51c765))
* improve SEO score of the landing page ([6ab0131](https://github.com/meysam81/tarzan/commit/6ab0131192b4a3164b22842186522a1d59dd6e04))
* improve the copy further, simplifying the language ([1a842ed](https://github.com/meysam81/tarzan/commit/1a842ed2e313e0e8a68c84a649cb03cd599774a7))
* increase SEO score to 100 :) ([2049ee8](https://github.com/meysam81/tarzan/commit/2049ee81063c94f3758428e53db031678ec1d129))
* record and persist CSP violations for later audit and troubleshooting ([75674c0](https://github.com/meysam81/tarzan/commit/75674c0275bbc28f48e042f323facb6c83f94a3d))
* save the webhook payload for later audit ([d1e9946](https://github.com/meysam81/tarzan/commit/d1e9946ab8496977867d013d859ab281e2820d8c))


### Bug Fixes

* add data to image source ([65bfb1a](https://github.com/meysam81/tarzan/commit/65bfb1a4b5a14e484b205c34ba94b93dd0d712d1))
* add the font and connect to CSP Header ([46b314e](https://github.com/meysam81/tarzan/commit/46b314e6561a5bf29a4f4d4d871d2e419d17a6b7))
* **build:** do not update go.sum on build ([87c7299](https://github.com/meysam81/tarzan/commit/87c7299dda35422f5e9211973448816e4d233387))
* **CI:** add gitkeep to dist ([3a548d8](https://github.com/meysam81/tarzan/commit/3a548d89c13e04b80c3e777739dbc8855e2a7f1e))
* **CI:** create dist in the job ([78dd411](https://github.com/meysam81/tarzan/commit/78dd4112e084221f07d97bf5cf9d66505784f38d))
* **CI:** make linter happy ([6151b12](https://github.com/meysam81/tarzan/commit/6151b12fa5fa349105f344d0ff3a5b6ff6a58df9))
* **CI:** make linter happy ([9b6b93b](https://github.com/meysam81/tarzan/commit/9b6b93b07872a30acfc75115ac2ae1291be208f3))
* **CI:** pass the e2e tests ([f0abcb5](https://github.com/meysam81/tarzan/commit/f0abcb5175d335b7c1efde76fcef826fcf1eceeb))
* **e2e:** assert 204 on webhook success ([57d6f1c](https://github.com/meysam81/tarzan/commit/57d6f1c3c54b51b9963c5d4f36c950bf5d07d333))
* handle images with both html and text body ([b52b0e2](https://github.com/meysam81/tarzan/commit/b52b0e263a236fdde940e650aaac3e1e2f3a8622))
* remove go back from 404 not found ([0a3549e](https://github.com/meysam81/tarzan/commit/0a3549e8e663d880a6b6722d379ba9c151115c28))
* serve the attachments from its own handler ([7bd78de](https://github.com/meysam81/tarzan/commit/7bd78de90b7d37da97dbe0b0299c14db2b5b356e))
* **server:** use html body for attachment rendering ([8e364b9](https://github.com/meysam81/tarzan/commit/8e364b9e63bcdc2dc3a9868853b33599e2ac0276))

## [1.0.2](https://github.com/meysam81/tarzan/compare/v1.0.1...v1.0.2) (2025-05-27)


### Bug Fixes

* **CI:** provide PAT to github releaser ([93e722b](https://github.com/meysam81/tarzan/commit/93e722b564e60e16378edb4835bdf9b2d211b014))

## [1.0.1](https://github.com/meysam81/tarzan/compare/v1.0.0...v1.0.1) (2025-05-27)


### Bug Fixes

* **CI:** build and publish compiled binaries to github assets ([e07e552](https://github.com/meysam81/tarzan/commit/e07e552d05be45d63c70f6bf9de48dd7eb40ee96))

## 1.0.0 (2025-05-27)


### Features

* add .env.example ([0dcca0c](https://github.com/meysam81/tarzan/commit/0dcca0cb7559599ba3ecdfce6b150f805e8e9fc4))
* add a title and styling for hover to it ([d1e8055](https://github.com/meysam81/tarzan/commit/d1e8055eaf4f3aeca8456f8c1be6e51db3ea2398))
* add clear instructions on usage at the hero section ([fe94ef6](https://github.com/meysam81/tarzan/commit/fe94ef62ae7269e5ff5141c6e0c45d8dece2051c))
* add csp validation report handler endpoint ([704c628](https://github.com/meysam81/tarzan/commit/704c628a4af60cab9be3425ffa827127a2bdf617))
* add post page navigation and remove the share story footer ([7f368b8](https://github.com/meysam81/tarzan/commit/7f368b8324810438d218c32b2c6956767ee22b2e))
* add PR template to github ([e5dc987](https://github.com/meysam81/tarzan/commit/e5dc987cd1163e6b1846144745e9645d01c531af))
* add prometheus exporter with basic auth configuration ([c219861](https://github.com/meysam81/tarzan/commit/c219861149886a85e2c3ecbe878a5c23393a6362))
* add redis datastore ([1a506d4](https://github.com/meysam81/tarzan/commit/1a506d48acb4329e6020aa288aba6585e48d2db5))
* add redis filestore ([9468f73](https://github.com/meysam81/tarzan/commit/9468f73769b40151c8d364ff0cfab7c387c574cf))
* add robots.txt to disallow everything ([efc8a27](https://github.com/meysam81/tarzan/commit/efc8a2717f056a2e1d439f82b14e3cba0a746d74))
* **CI:** add lychee config ([2638778](https://github.com/meysam81/tarzan/commit/2638778bcd0ca0e6a55c330ea077d2ed76e1614c))
* **CI:** add lychee job ([7baf45c](https://github.com/meysam81/tarzan/commit/7baf45c4c9cdb1b4edeeb30ecdfd462edb2302b6))
* **CI:** build and publish docker to dockerhub ([c6b109f](https://github.com/meysam81/tarzan/commit/c6b109ff534e3aabd5eb6ed8dffee7fa77ed23e6))
* **config:** add base url and remove frontend ([f261a62](https://github.com/meysam81/tarzan/commit/f261a62e6f65a5c0a22d52e535968056208872d0))
* create blog post view page ([33fe090](https://github.com/meysam81/tarzan/commit/33fe0901439d626e90e191228d71d514644407b3))
* create fat CGO disabled binary ([504260c](https://github.com/meysam81/tarzan/commit/504260cfcf4a2305c1f6733956ce07158be10b06))
* create filestore as interface to allow redis implementation ([de582c7](https://github.com/meysam81/tarzan/commit/de582c787e86a4ed2ce8a34edc827d195ef32347))
* create frontend code structure & serve by go ([f69bb0c](https://github.com/meysam81/tarzan/commit/f69bb0c6502856cfb15c612792da71cec6172d89))
* create the repo structure with CI and docker ([fd19a88](https://github.com/meysam81/tarzan/commit/fd19a888cfa9715a2d003eda55b1428fb34f1576))
* **deploy:** add k8s manifests ([fcf3579](https://github.com/meysam81/tarzan/commit/fcf3579ecd8444515f0a6d02cacf426bf2d908e8))
* **docs:** add k8s to the deploy section ([77249f6](https://github.com/meysam81/tarzan/commit/77249f67a2b2efed9d41ddb82d6f794919a41d68))
* finish up the UI with an exceptional style ([a0a6f49](https://github.com/meysam81/tarzan/commit/a0a6f499b9a84c86ab47d3a6c6056583959ffc9e))
* implement robust client-server websocket connection for live updates ([8060dc5](https://github.com/meysam81/tarzan/commit/8060dc5a6361dbe15f73a18aa09918479a8adfaa))
* increase extensibility when creating datastore ([63479b2](https://github.com/meysam81/tarzan/commit/63479b2b204980f80153615e57e7edb4cfa31fb0))
* move authorized senders to the database ([6b1e76d](https://github.com/meysam81/tarzan/commit/6b1e76dfc54a280a8e6c527afbb2c8e28a7ab575))
* overhault rewrite of the index.html and cleanup configs ([e8ae7af](https://github.com/meysam81/tarzan/commit/e8ae7afa9e78db40b40baf324e148b9d680bea03))
* receive credentials from env/config ([c5b2e03](https://github.com/meysam81/tarzan/commit/c5b2e035f4f2cb51b86f633a4e25318d366a9329))
* redesign the UI entirely ([b6651f9](https://github.com/meysam81/tarzan/commit/b6651f95b46ef53317d884766a06c332be141a18))
* render author name with tooltip of email ([b79559b](https://github.com/meysam81/tarzan/commit/b79559b6513625313f8340b13af8ea73a4c80829))
* rewrite the frontend in vuejs ([0e15383](https://github.com/meysam81/tarzan/commit/0e15383a5cde274d070ed6b953d475cde77a54b0))
* **server:** create websocket for live post update ([f078a2e](https://github.com/meysam81/tarzan/commit/f078a2e651c22775dbab31a09e0004b008e93af2))
* **server:** rewrite the router to use chi framework ([c9736bd](https://github.com/meysam81/tarzan/commit/c9736bd5642b3eecff2d02f66c2a603d72aef647))
* store file attachments in redis filestore ([773eba0](https://github.com/meysam81/tarzan/commit/773eba0db5fd8364697d7b0196bd5471f69f1d4f))
* thoroughly rewrite the copy in favor of clear messaging ([2221a2e](https://github.com/meysam81/tarzan/commit/2221a2ec9d665df1f211df1b78e17007c912c0c6))


### Bug Fixes

* add unsafe inline to script source ([d4d8cd1](https://github.com/meysam81/tarzan/commit/d4d8cd18642178294475e3f890f6ae7c5460b039))
* **build:** add CA certificates to the final docker file ([d14bdf4](https://github.com/meysam81/tarzan/commit/d14bdf4c6ed9f500b0abb4edf0cc099c694920b4))
* **build:** modify docker ignore to include the required paths ([8d9c163](https://github.com/meysam81/tarzan/commit/8d9c163593288277c2e422d5163a74eb17521f0d))
* **build:** remove build failure on missing dist ([28a7805](https://github.com/meysam81/tarzan/commit/28a7805bc903fb56a5eb7e2bd520c2a960ff0c6d))
* **CI:** allow creating issues ([2ef0638](https://github.com/meysam81/tarzan/commit/2ef063862c5cbb8c4b1d071e6245738f6f13f488))
* **CI:** change release-please language to go ([a4be446](https://github.com/meysam81/tarzan/commit/a4be4465eb3dda74bf55a05f7cb4803d371f6859))
* **CI:** make linter happy ([6a49663](https://github.com/meysam81/tarzan/commit/6a49663527a98953394f8fa88370d657d7f528d1))
* **CI:** make linter happy ([3fdfb7a](https://github.com/meysam81/tarzan/commit/3fdfb7a9bfaed8e1d0c3648d30d979591e577aca))
* **CI:** make linter happy ([1f8915d](https://github.com/meysam81/tarzan/commit/1f8915d75ff06f072cad53307312c272fb639ea0))
* **CI:** make linter happy ([3a799d9](https://github.com/meysam81/tarzan/commit/3a799d9f38776536cf1cc3a7f6d190b7dfe96cfa))
* **CI:** make linter happy, again ([155a0e5](https://github.com/meysam81/tarzan/commit/155a0e520ebde0d4a87918f9734e3e233d56f091))
* **CI:** make linter happy, again ([9a57676](https://github.com/meysam81/tarzan/commit/9a57676cb3f96a86b249872a6fef5b177eb3697a))
* dd unsafe inline to style source ([cdf6f27](https://github.com/meysam81/tarzan/commit/cdf6f27462b79ac8a822043cb32d45b65183c3e5))
* **deploy:** use secure websocket for prod ([6adacfe](https://github.com/meysam81/tarzan/commit/6adacfe83ad741bd5fa3d2772fc3776310423c76))
* make linter happy ([4a32eee](https://github.com/meysam81/tarzan/commit/4a32eee131dbf9bf3f5d7a934366f6480a3e814d))
* modify CSP header that broke scripts ([c980d06](https://github.com/meysam81/tarzan/commit/c980d064312c0d849b99bce3d7f5d80f923cd6b5))
* parse the date with the correct format ([2f5c7a5](https://github.com/meysam81/tarzan/commit/2f5c7a53660e4d277cadf317ada83746d01487ee))
* put PR template in its own dir ([72bf161](https://github.com/meysam81/tarzan/commit/72bf16155eae40dce2b1b8be2b541688479c73e3))
* quote the unsafe inline ([c27ebf5](https://github.com/meysam81/tarzan/commit/c27ebf5a8c507f74546d4cc224af653d1cc26cff))
* record metrics at the end ([8c5b9bb](https://github.com/meysam81/tarzan/commit/8c5b9bb971e255d8276f52473b9e03252a629860))
* remove that absurd animated text from the main screen ([f99ea69](https://github.com/meysam81/tarzan/commit/f99ea693e405b35c48c397efb09bf9e22da4ccdc))
* remove the absolute vs relative path nonsense ([43c307b](https://github.com/meysam81/tarzan/commit/43c307bb589c8da4b3a44fa0c5edee43fe553fc7))
* serve statics from the root path ([1f042b9](https://github.com/meysam81/tarzan/commit/1f042b96e2607325f559fb4566ce528dfed131e3))
* specify the path in the docker env ([4e0f6fb](https://github.com/meysam81/tarzan/commit/4e0f6fbd1adaa7a9b1958708ac6054b72298edba))
* specify websocket url at frontend build time ([d82d971](https://github.com/meysam81/tarzan/commit/d82d97148335e64a009ca2f6521d2d6a8ea18d13))
* **style:** change the contrast setting of the hero text ([bc7b2f2](https://github.com/meysam81/tarzan/commit/bc7b2f2b8ba9de463c458a33944c9246d6f9bc45))
* **UI:** render the how it works modal correctly ([55fa9e3](https://github.com/meysam81/tarzan/commit/55fa9e332c1fe085b267c98bcf0731bbeabf45bc))
* **UI:** use ref for accessing env var ([c4daac4](https://github.com/meysam81/tarzan/commit/c4daac417b6f09b4a0d1e3cc25ab1d65be966065))
* update bun lock file ([56613b7](https://github.com/meysam81/tarzan/commit/56613b734493e44a0a2dd84e84b4b5a5075315f1))
* update deps and drop local replace ([0fa0030](https://github.com/meysam81/tarzan/commit/0fa00308d5ea8ba98b6067663354899e92070c90))
* update lock file ([f799884](https://github.com/meysam81/tarzan/commit/f799884217240f88e44561c7eb26efbc00ade957))
