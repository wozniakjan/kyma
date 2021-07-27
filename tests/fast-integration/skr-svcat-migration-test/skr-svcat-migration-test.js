const uuid = require("uuid");
const { 
  KEBConfig,
  KEBClient,
  provisionSKR,
  deprovisionSKR,
} = require("../kyma-environment-broker");
const {
  GardenerConfig,
  GardenerClient
} = require("../gardener");
const {
  debug,
  genRandom,
  getEnvOrThrow,
  initializeK8sClient,
} = require("../utils");
const t = require("./test-helpers");
const fs = require('fs');
const os = require('os');

describe("SKR SVCAT migration test", function() {
  const keb = new KEBClient(KEBConfig.fromEnv());
  const gardener = new GardenerClient(GardenerConfig.fromEnv());

  const suffix = genRandom(4);
  const appName = `app-${suffix}`;
  const runtimeName = `kyma-${suffix}`;
  const runtimeID = uuid.v4();
  
  const svcatPlatform = `svcat-${suffix}`
  const btpOperatorInstance = `btp-operator-${suffix}`
  const btpOperatorBinding = `btp-operator-binding-${suffix}`

  debug(`RuntimeID ${runtimeID}`, `Runtime ${runtimeName}`, `Application ${appName}`, `Suffix ${suffix}`);

  const testNS = "skr-test";
  const AWS_PLAN_ID = "361c511f-f939-4621-b228-d0fb79a1fe15";

  const clientID = getEnvOrThrow("BTP_OPERATOR_CLIENTID");
  const clientSecret = getEnvOrThrow("BTP_OPERATOR_CLIENTSECRET");
  const URL = getEnvOrThrow("BTP_OPERATOR_URL");

  this.timeout(60 * 60 * 1000 * 3); // 3h
  this.slow(5000);  

  let skr;
  it(`Should provision SKR`, async function() {
    fs.mkdirSync(`${os.homedir()}/.kube`, true);
    skr = await provisionSKR(keb, gardener, runtimeID, runtimeName);
    fs.writeFileSync(`${os.homedir()}/.kube/config`, skr.shoot.kubeconfig);
  });
  it(`Should initialize K8s`, async function() {
    await initializeK8sClient({kubeconfig: skr.shoot.kubeconfig});
  });
  let creds;
  it(`Should instantiate SM Instance and Binding`, async function() {
    creds = await t.smInstanceBinding(URL, clientID, clientSecret, svcatPlatform, btpOperatorInstance, btpOperatorBinding);
  });
  it(`Should install helm charts`, async function() {
    await t.installHelmCharts(creds);
  });
  it(`Should deprovision SKR`, async function() {
    await deprovisionSKR(keb, runtimeID);
  });
  it(`Should cleanup SM instances and bindings`, async function() {
    await t.cleanupInstanceBinding(URL, clientID, clientSecret, svcatPlatform, btpOperatorInstance, btpOperatorBinding);
  });
});
