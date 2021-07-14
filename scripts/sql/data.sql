INSERT INTO `baetyl_property` (`name`, `value`) VALUES
('sync-server-address', 'https://host.docker.internal:9005'),
('init-server-address', 'https://host.docker.internal:9003'),

('command-docker-installation', 'curl -sSL https://get.daocloud.io/docker | sh'),
('command-k3s-installation-containerd', 'curl -sfL http://rancher-mirror.cnrancher.com/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn INSTALL_K3S_VERSION=v1.18.9+k3s1 INSTALL_K3S_EXEC=\"--write-kubeconfig ~/.kube/config --write-kubeconfig-mode 666\" sh -'),
('command-k3s-installation-docker', 'curl -sfL http://rancher-mirror.cnrancher.com/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn INSTALL_K3S_VERSION=v1.18.9+k3s1 INSTALL_K3S_EXEC=\"--docker --write-kubeconfig ~/.kube/config --write-kubeconfig-mode 666\" sh -'),
('baetyl-init-command', 'curl -skfL \'{{GetProperty \"init-server-address\"}}/v1/init/baetyl-install.sh?token={{.Token}}&mode={{.mode}}&initApplyYaml={{.InitApplyYaml}}\' -osetup.sh && sh setup.sh'),
('command-baetyl-kube-delete', 'kubectl delete ns baetyl-edge baetyl-edge-system --grace-period=0 --force'),
('command-baetyl-native-delete', 'sudo baetyl delete && sudo baetyl delete -n baetyl-edge'),

('command-baetyl-installation', 'curl -sSL https://baetyl-repo-gz.gz.bcebos.com/v2/install.sh | bash');

INSERT INTO `baetyl_module` (`name`, `version`, `image`, `programs`, `type`, `is_latest`, `description`) VALUES
('baetyl-broker', 'v2.1.0', 'hub.baidubce.com/baetyl/broker:v2.1.0', '{\"darwin-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/baetyl-broker_darwin-amd64_v2.1.0.zip\",\"linux-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/baetyl-broker_linux-amd64_v2.1.0.zip\",\"linux-arm64-v8\":\"https://baetyl.bj.bcebos.com/test/programs/baetyl-broker_linux-arm64_v2.1.0.zip\",\"linux-arm-v7\":\"https://baetyl.bj.bcebos.com/test/programs/baetyl-broker_linux-arm-v7_v2.1.0.zip\"}', 'system', 1, 'for baetyl-broker'),
('baetyl', 'v2.2.2', 'docker.io/baetyltech/baetyl:v2.2.2', '{\"darwin-amd64\":\"https://baetyl.bj.bcebos.com/baetyl-2.2.2/baetyl_darwin-amd64_v2.2.2.zip\",\"linux-amd64\":\"https://baetyl.bj.bcebos.com/baetyl-2.2.2/baetyl_linux-amd64_v2.2.2.zip\",\"linux-arm64-v8\":\"https://baetyl.bj.bcebos.com/baetyl-2.2.2/baetyl_linux-arm64_v2.2.2.zip\",\"linux-arm-v7\":\"https://baetyl.bj.bcebos.com/baetyl-2.2.2/baetyl_linux-arm-v7_v2.2.2.zip\"}', 'system', 1, 'for baetyl'),
('sql', 'v2.1.1', 'hub.baidubce.com/baetyl/sql:v2.1.1', '{\"darwin-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/sql/baetyl-sql_darwin-amd64_git-c69c83e.zip\",\"linux-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/sql/baetyl-sql_linux-amd64_git-c69c83e.zip\",\"linux-arm64-v8\":\"https://baetyl.bj.bcebos.com/test/programs/sql/baetyl-sql_linux-arm64_git-c69c83e.zip\",\"linux-arm-v7\":\"https://baetyl.bj.bcebos.com/test/programs/sql/baetyl-sql_linux-arm-v7_git-c69c83e.zip\"}', 'runtime_user', 1, 'for sql'),
('baetyl-function', 'v2.1.1', 'hub.baidubce.com/baetyl/function:v2.1.1', '{\"darwin-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/function/baetyl-function_darwin-amd64_git-2c804da.zip\",\"linux-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/function/baetyl-function_linux-amd64_git-2c804da.zip\",\"linux-arm64-v8\":\"https://baetyl.bj.bcebos.com/test/programs/function/baetyl-function_linux-arm64_git-2c804da.zip\",\"linux-arm-v7\":\"https://baetyl.bj.bcebos.com/test/programs/function/baetyl-function_linux-arm-v7_git-2c804da.zip\"}', 'opt_system', 1, '该应用是 baetyl 框架端侧的函数计算框架，支持前端代理和后端函数运行模块的请求传输。'),
('baetyl-rule', 'v2.1.1', 'hub.baidubce.com/baetyl/rule:v2.1.1', '{\"darwin-amd64\":\"https://baetyl.cdn.bcebos.com/test/programs/rule/baetyl-rule_darwin-amd64_git-8d7f716.zip\",\"linux-amd64\":\"https://baetyl.cdn.bcebos.com/test/programs/rule/baetyl-rule_linux-amd64_git-8d7f716.zip\",\"linux-arm64-v8\":\"https://baetyl.cdn.bcebos.com/test/programs/rule/baetyl-rule_linux-arm64_git-8d7f716.zip\",\"linux-arm-v7\":\"https://baetyl.cdn.bcebos.com/test/programs/rule/baetyl-rule_linux-arm64_git-8d7f716.zip\"}', 'opt_system', 1, '该应用实现 baetyl 框架端侧的消息流转和交换。'),
('nodejs10', 'v2.1.1', 'hub.baidubce.com/baetyl/nodejs10:3.6-v2.1.1', '{\"darwin-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/nodejs10/new/baetyl-nodejs10_darwin-amd64_git-2c804da.zip\",\"linux-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/nodejs10/new/baetyl-nodejs10_linux-amd64_git-2c804da.zip\",\"linux-arm64-v8\":\"https://baetyl.bj.bcebos.com/test/programs/nodejs10/new/baetyl-nodejs10_linux-arm64_git-2c804da.zip\",\"linux-arm-v7\":\"https://baetyl.bj.bcebos.com/test/programs/nodejs10/new/baetyl-nodejs10_linux-arm-v7_git-2c804da.zip\"}', 'runtime_user', 1, 'for nodejs10'),
('python3', 'v2.1.1', 'hub.baidubce.com/baetyl/python3:3.6-v2.1.1', '{\"darwin-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/python3/new/baetyl-python3_darwin-amd64_git-2c804da.zip\",\"linux-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/python3/new/baetyl-python3_linux-amd64_git-2c804da.zip\",\"linux-arm64-v8\":\"https://baetyl.bj.bcebos.com/test/programs/python3/new/baetyl-python3_linux-arm64_git-2c804da.zip\",\"linux-arm-v7\":\"https://baetyl.bj.bcebos.com/test/programs/python3/new/baetyl-python3_linux-arm-v7_git-2c804da.zip\"}', 'runtime_user', 1, 'for python3'),
('python3-opencv', 'v2.1.1', 'hub.baidubce.com/baetyl-sandbox/python3-opencv:3.6-v2.1.1', '{\"darwin-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/python3/new/baetyl-python3_darwin-amd64_git-2c804da.zip\",\"linux-amd64\":\"https://baetyl.bj.bcebos.com/test/programs/python3/new/baetyl-python3_linux-amd64_git-2c804da.zip\",\"linux-arm-v7\":\"https://baetyl.bj.bcebos.com/test/programs/python3/new/baetyl-python3_linux-arm-v7_git-2c804da.zip\",\"linux-arm64-v8\":\"https://baetyl.bj.bcebos.com/test/programs/python3/new/baetyl-python3_linux-arm64_git-2c804da.zip\"}', 'runtime_user', 1, 'for python3-opencv df');