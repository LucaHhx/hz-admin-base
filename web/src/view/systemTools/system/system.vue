<template>
  <div class="system">
    <el-form
      ref="form"
      :model="config"
      label-width="240px"
    >
      <!--  System start  -->
      <el-tabs v-model="activeNames">
        <el-tab-pane
          :label="$t('system.title')"
          name="1"
          class="mt-3.5"
        >
          <el-form-item :label="$t('system.port')">
            <el-input-number
              v-model="config.system.addr"
              :placeholder="$t('system.portPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.dbType')">
            <el-select
              v-model="config.system['db-type']"
              class="w-full"
            >
              <el-option value="mysql" />
              <el-option value="pgsql" />
              <el-option value="mssql" />
              <el-option value="sqlite" />
              <el-option value="oracle" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('system.ossType')">
            <el-select
              v-model="config.system['oss-type']"
              class="w-full"
            >
              <el-option
                :value="'local'"
                :label="$t('system.local')"
              />
              <el-option
                :value="'qiniu'"
                :label="$t('system.qiniu')"
              />
              <el-option
                :value="'tencent-cos'"
                :label="$t('system.tencentCos')"
              />
              <el-option
                :value="'aliyun-oss'"
                :label="$t('system.aliyunOss')"
              />
              <el-option
                :value="'huawei-obs'"
                :label="$t('system.huaweiObs')"
              />
              <el-option
                :value="'cloudflare-r2'"
                :label="$t('system.cloudflareR2')"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('system.multipoint')">
            <el-switch v-model="config.system['use-multipoint']" />
          </el-form-item>
          <el-form-item :label="$t('system.useRedis')">
            <el-switch v-model="config.system['use-redis']" />
          </el-form-item>
          <el-form-item :label="$t('system.useMongo')">
            <el-switch v-model="config.system['use-mongo']" />
          </el-form-item>
          <el-form-item :label="$t('system.strictAuth')">
            <el-switch v-model="config.system['use-strict-auth']" />
          </el-form-item>
          <el-form-item :label="$t('system.ipLimitCount')">
            <el-input-number v-model.number="config.system['iplimit-count']" />
          </el-form-item>
          <el-form-item :label="$t('system.ipLimitTime')">
            <el-input-number v-model.number="config.system['iplimit-time']" />
          </el-form-item>
          <el-tooltip
            :content="$t('system.routerPrefixTip')"
            placement="top-start"
          >
            <el-form-item :label="$t('system.routerPrefix')">
              <el-input
                v-model.trim="config.system['router-prefix']"
                :placeholder="$t('system.routerPrefixPlaceholder')"
              />
            </el-form-item>
          </el-tooltip>
        </el-tab-pane>
        <el-tab-pane
          :label="$t('system.jwt.title')"
          name="2"
          class="mt-3.5"
        >
          <el-form-item :label="$t('system.jwt.signingKey')">
            <el-input
              v-model.trim="config.jwt['signing-key']"
              :placeholder="$t('system.jwt.signingKeyPlaceholder')"
            >
              <template #append>
                <el-button @click="getUUID">
                  {{ $t('system.jwt.generate') }}
                </el-button>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item :label="$t('system.jwt.expiresTime')">
            <el-input
              v-model.trim="config.jwt['expires-time']"
              :placeholder="$t('system.jwt.expiresTimePlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.jwt.bufferTime')">
            <el-input
              v-model.trim="config.jwt['buffer-time']"
              :placeholder="$t('system.jwt.bufferTimePlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.jwt.issuer')">
            <el-input
              v-model.trim="config.jwt.issuer"
              :placeholder="$t('system.jwt.issuerPlaceholder')"
            />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
          :label="$t('system.zap.title')"
          name="3"
          class="mt-3.5"
        >
          <el-form-item :label="$t('system.zap.level')">
            <el-select v-model="config.zap.level">
              <el-option
                value="off"
                :label="$t('system.zap.levelOptions.off')"
              />
              <el-option
                value="fatal"
                :label="$t('system.zap.levelOptions.fatal')"
              />
              <el-option
                value="error"
                :label="$t('system.zap.levelOptions.error')"
              />
              <el-option
                value="warn"
                :label="$t('system.zap.levelOptions.warn')"
              />
              <el-option
                value="info"
                :label="$t('system.zap.levelOptions.info')"
              />
              <el-option
                value="debug"
                :label="$t('system.zap.levelOptions.debug')"
              />
              <el-option
                value="trace"
                :label="$t('system.zap.levelOptions.trace')"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('system.zap.format')">
            <el-select v-model="config.zap.format">
              <el-option
                value="console"
                label="console"
              />
              <el-option
                value="json"
                label="json"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('system.zap.prefix')">
            <el-input
              v-model.trim="config.zap.prefix"
              :placeholder="$t('system.zap.prefixPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.zap.director')">
            <el-input
              v-model.trim="config.zap.director"
              :placeholder="$t('system.zap.directorPlaceholder')"
            />
          </el-form-item>
          <el-form-item
            :label="$t('system.zap.encodeLevel')"
            class="w-6/12"
          >
            <el-select v-model="config.zap['encode-level']">
              <el-option
                value="LowercaseLevelEncoder"
                label="LowercaseLevelEncoder"
              />
              <el-option
                value="LowercaseColorLevelEncoder"
                label="LowercaseColorLevelEncoder"
              />
              <el-option
                value="CapitalLevelEncoder"
                label="CapitalLevelEncoder"
              />
              <el-option
                value="CapitalColorLevelEncoder"
                label="CapitalColorLevelEncoder"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('system.zap.stacktraceKey')">
            <el-input
              v-model.trim="config.zap['stacktrace-key']"
              :placeholder="$t('system.zap.stacktraceKeyPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.zap.retentionDay')">
            <el-input-number v-model="config.zap['retention-day']" />
          </el-form-item>
          <el-form-item :label="$t('system.zap.showLine')">
            <el-switch v-model="config.zap['show-line']" />
          </el-form-item>
          <el-form-item :label="$t('system.zap.logInConsole')">
            <el-switch v-model="config.zap['log-in-console']" />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
          v-if="config.system['use-redis']"
          :label="$t('system.redisConfig.title')"
          name="4"
          class="mt-3.5"
        >
          <el-form-item :label="$t('system.redisConfig.db')">
            <el-input-number
              v-model="config.redis.db"
              min="0"
              max="16"
            />
          </el-form-item>
          <el-form-item :label="$t('system.redisConfig.addr')">
            <el-input
              v-model.trim="config.redis.addr"
              :placeholder="$t('system.redisConfig.addrPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.redisConfig.password')">
            <el-input
              v-model.trim="config.redis.password"
              :placeholder="$t('system.redisConfig.passwordPlaceholder')"
            />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
          :label="$t('system.email.title')"
          name="5"
          class="mt-3.5"
        >
          <el-form-item :label="$t('system.email.to')">
            <el-input
              v-model="config.email.to"
              :placeholder="$t('system.email.toPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.email.port')">
            <el-input-number v-model="config.email.port" />
          </el-form-item>
          <el-form-item :label="$t('system.email.from')">
            <el-input
              v-model.trim="config.email.from"
              :placeholder="$t('system.email.fromPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.email.host')">
            <el-input
              v-model.trim="config.email.host"
              :placeholder="$t('system.email.hostPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.email.ssl')">
            <el-switch v-model="config.email['is-ssl']" />
          </el-form-item>
          <el-form-item :label="$t('system.email.secret')">
            <el-input
              v-model.trim="config.email.secret"
              :placeholder="$t('system.email.secretPlaceholder')"
            />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
          v-if="config.system['use-mongo']"
          :label="$t('system.mongoConfig.title')"
          name="14"
          class="mt-3.5"
        >
          <el-form-item :label="$t('system.mongoConfig.coll')">
            <el-input
              v-model.trim="config.mongo.coll"
              :placeholder="$t('system.mongoConfig.collPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.mongoConfig.options')">
            <el-input
              v-model.trim="config.mongo.options"
              :placeholder="$t('system.mongoConfig.optionsPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.mongoConfig.database')">
            <el-input
              v-model.trim="config.mongo.database"
              :placeholder="$t('system.mongoConfig.databasePlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.mongoConfig.username')">
            <el-input
              v-model.trim="config.mongo.username"
              :placeholder="$t('system.mongoConfig.usernamePlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.mongoConfig.password')">
            <el-input
              v-model.trim="config.mongo.password"
              :placeholder="$t('system.mongoConfig.passwordPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.mongoConfig.minPoolSize')">
            <el-input-number
              v-model="config.mongo['min-pool-size']"
              min="0"
            />
          </el-form-item>
          <el-form-item :label="$t('system.mongoConfig.maxPoolSize')">
            <el-input-number
              v-model="config.mongo['max-pool-size']"
              min="100"
            />
          </el-form-item>
          <el-form-item :label="$t('system.mongoConfig.socketTimeoutMs')">
            <el-input-number
              v-model="config.mongo['socket-timeout-ms']"
              min="0"
            />
          </el-form-item>
          <el-form-item :label="$t('system.mongoConfig.connectTimeoutMs')">
            <el-input-number
              v-model="config.mongo['connect-timeout-ms']"
              min="0"
            />
          </el-form-item>
          <el-form-item :label="$t('system.mongoConfig.isZap')">
            <el-switch v-model="config.mongo['is-zap']" />
          </el-form-item>
          <el-form-item
            v-for="(item, k) in config.mongo.hosts"
            :key="k"
            :label="$t('system.mongoConfig.node', { index: k + 1 })"
          >
            <div
              v-for="(_, k2) in item"
              :key="k2"
            >
              <el-form-item
                :key="k + k2"
                :label="k2"
                label-width="60"
              >
                <el-input
                  v-model.trim="item[k2]"
                  :placeholder="k2 === 'host' ? $t('system.mongoConfig.hostPlaceholder') : $t('system.mongoConfig.portPlaceholder')"
                />
              </el-form-item>
            </div>
            <el-form-item v-if="k > 0">
              <el-button
                type="danger"
                size="small"
                plain
                :icon="Minus"
                class="ml-3"
                @click="removeNode(k)"
              />
            </el-form-item>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              size="small"
              plain
              :icon="Plus"
              @click="addNode"
            />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
          :label="$t('system.captcha.title')"
          name="7"
          class="mt-3.5"
        >
          <el-form-item :label="$t('system.captcha.keyLong')">
            <el-input-number
              v-model="config.captcha['key-long']"
              :min="4"
              :max="6"
            />
          </el-form-item>
          <el-form-item :label="$t('system.captcha.imgWidth')">
            <el-input-number v-model.number="config.captcha['img-width']" />
          </el-form-item>
          <el-form-item :label="$t('system.captcha.imgHeight')">
            <el-input-number v-model.number="config.captcha['img-height']" />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
          :label="$t('system.dbConfig.title')"
          name="9"
          class="mt-3.5"
        >
          <template v-if="config.system['db-type'] === 'mysql'">
            <el-form-item label="">
              <h3>{{ $t('system.dbConfig.mysql.title') }}</h3>
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.username')">
              <el-input
                v-model.trim="config.mysql.username"
                :placeholder="$t('system.dbConfig.mysql.usernamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.password')">
              <el-input
                v-model.trim="config.mysql.password"
                :placeholder="$t('system.dbConfig.mysql.passwordPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.path')">
              <el-input
                v-model.trim="config.mysql.path"
                :placeholder="$t('system.dbConfig.mysql.pathPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.dbName')">
              <el-input
                v-model.trim="config.mysql['db-name']"
                :placeholder="$t('system.dbConfig.mysql.dbNamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.prefix')">
              <el-input
                v-model.trim="config.mysql['prefix']"
                :placeholder="$t('system.dbConfig.mysql.prefixPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.singular')">
              <el-switch v-model="config.mysql['singular']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.engine')">
              <el-input
                v-model.trim="config.mysql['engine']"
                :placeholder="$t('system.dbConfig.mysql.enginePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.maxIdleConns')">
              <el-input-number
                v-model="config.mysql['max-idle-conns']"
                :min="1"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.maxOpenConns')">
              <el-input-number
                v-model="config.mysql['max-open-conns']"
                :min="1"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.logZap')">
              <el-switch v-model="config.mysql['log-zap']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mysql.logMode')">
              <el-select v-model="config.mysql['log-mode']">
                <el-option
                  :value="'off'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.off')"
                />
                <el-option
                  :value="'fatal'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.fatal')"
                />
                <el-option
                  :value="'error'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.error')"
                />
                <el-option
                  :value="'warn'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.warn')"
                />
                <el-option
                  :value="'info'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.info')"
                />
                <el-option
                  :value="'debug'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.debug')"
                />
                <el-option
                  :value="'trace'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.trace')"
                />
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'pgsql'">
            <el-form-item label="">
              <h3>{{ $t('system.dbConfig.pgsql.title') }}</h3>
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.username')">
              <el-input
                v-model="config.pgsql.username"
                :placeholder="$t('system.dbConfig.pgsql.usernamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.password')">
              <el-input
                v-model="config.pgsql.password"
                :placeholder="$t('system.dbConfig.pgsql.passwordPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.path')">
              <el-input
                v-model.trim="config.pgsql.path"
                :placeholder="$t('system.dbConfig.pgsql.pathPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.dbName')">
              <el-input
                v-model.trim="config.pgsql['db-name']"
                :placeholder="$t('system.dbConfig.pgsql.dbNamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.prefix')">
              <el-input
                v-model.trim="config.pgsql['prefix']"
                :placeholder="$t('system.dbConfig.pgsql.prefixPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.singular')">
              <el-switch v-model="config.pgsql['singular']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.engine')">
              <el-input
                v-model.trim="config.pgsql['engine']"
                :placeholder="$t('system.dbConfig.pgsql.enginePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.maxIdleConns')">
              <el-input-number v-model="config.pgsql['max-idle-conns']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.maxOpenConns')">
              <el-input-number v-model="config.pgsql['max-open-conns']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.logZap')">
              <el-switch v-model="config.pgsql['log-zap']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.pgsql.logMode')">
              <el-select v-model="config.pgsql['log-mode']">
                <el-option
                  :value="'off'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.off')"
                />
                <el-option
                  :value="'fatal'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.fatal')"
                />
                <el-option
                  :value="'error'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.error')"
                />
                <el-option
                  :value="'warn'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.warn')"
                />
                <el-option
                  :value="'info'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.info')"
                />
                <el-option
                  :value="'debug'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.debug')"
                />
                <el-option
                  :value="'trace'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.trace')"
                />
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'mssql'">
            <el-form-item label="">
              <h3>{{ $t('system.dbConfig.mssql.title') }}</h3>
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.username')">
              <el-input
                v-model.trim="config.mssql.username"
                :placeholder="$t('system.dbConfig.mssql.usernamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.password')">
              <el-input
                v-model.trim="config.mssql.password"
                :placeholder="$t('system.dbConfig.mssql.passwordPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.path')">
              <el-input
                v-model.trim="config.mssql.path"
                :placeholder="$t('system.dbConfig.mssql.pathPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.port')">
              <el-input
                v-model.trim="config.mssql.port"
                :placeholder="$t('system.dbConfig.mssql.portPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.dbName')">
              <el-input
                v-model.trim="config.mssql['db-name']"
                :placeholder="$t('system.dbConfig.mssql.dbNamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.prefix')">
              <el-input
                v-model.trim="config.mssql['prefix']"
                :placeholder="$t('system.dbConfig.mssql.prefixPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.singular')">
              <el-switch v-model="config.mssql['singular']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.engine')">
              <el-input
                v-model.trim="config.mssql['engine']"
                :placeholder="$t('system.dbConfig.mssql.enginePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.maxIdleConns')">
              <el-input-number v-model="config.mssql['max-idle-conns']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.maxOpenConns')">
              <el-input-number v-model="config.mssql['max-open-conns']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.logZap')">
              <el-switch v-model="config.mssql['log-zap']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.mssql.logMode')">
              <el-select v-model="config.mssql['log-mode']">
                <el-option
                  :value="'off'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.off')"
                />
                <el-option
                  :value="'fatal'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.fatal')"
                />
                <el-option
                  :value="'error'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.error')"
                />
                <el-option
                  :value="'warn'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.warn')"
                />
                <el-option
                  :value="'info'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.info')"
                />
                <el-option
                  :value="'debug'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.debug')"
                />
                <el-option
                  :value="'trace'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.trace')"
                />
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'sqlite'">
            <el-form-item label="">
              <h3>{{ $t('system.dbConfig.sqlite.title') }}</h3>
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.sqlite.username')">
              <el-input
                v-model.trim="config.sqlite.username"
                :placeholder="$t('system.dbConfig.sqlite.usernamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.sqlite.password')">
              <el-input
                v-model.trim="config.sqlite.password"
                :placeholder="$t('system.dbConfig.sqlite.passwordPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.sqlite.path')">
              <el-input
                v-model.trim="config.sqlite.path"
                :placeholder="$t('system.dbConfig.sqlite.pathPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.sqlite.port')">
              <el-input
                v-model.trim="config.sqlite.port"
                :placeholder="$t('system.dbConfig.sqlite.portPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.sqlite.dbName')">
              <el-input
                v-model.trim="config.sqlite['db-name']"
                :placeholder="$t('system.dbConfig.sqlite.dbNamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.sqlite.maxIdleConns')">
              <el-input-number v-model="config.sqlite['max-idle-conns']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.sqlite.maxOpenConns')">
              <el-input-number v-model="config.sqlite['max-open-conns']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.sqlite.logZap')">
              <el-switch v-model="config.sqlite['log-zap']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.sqlite.logMode')">
              <el-select v-model="config.sqlite['log-mode']">
                <el-option
                  :value="'off'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.off')"
                />
                <el-option
                  :value="'fatal'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.fatal')"
                />
                <el-option
                  :value="'error'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.error')"
                />
                <el-option
                  :value="'warn'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.warn')"
                />
                <el-option
                  :value="'info'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.info')"
                />
                <el-option
                  :value="'debug'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.debug')"
                />
                <el-option
                  :value="'trace'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.trace')"
                />
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'oracle'">
            <el-form-item label="">
              <h3>{{ $t('system.dbConfig.oracle.title') }}</h3>
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.username')">
              <el-input
                v-model.trim="config.oracle.username"
                :placeholder="$t('system.dbConfig.oracle.usernamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.password')">
              <el-input
                v-model.trim="config.oracle.password"
                :placeholder="$t('system.dbConfig.oracle.passwordPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.path')">
              <el-input
                v-model.trim="config.oracle.path"
                :placeholder="$t('system.dbConfig.oracle.pathPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.dbName')">
              <el-input
                v-model.trim="config.oracle['db-name']"
                :placeholder="$t('system.dbConfig.oracle.dbNamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.prefix')">
              <el-input
                v-model.trim="config.oracle['prefix']"
                :placeholder="$t('system.dbConfig.oracle.prefixPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.singular')">
              <el-switch v-model="config.oracle['singular']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.engine')">
              <el-input
                v-model.trim="config.oracle['engine']"
                :placeholder="$t('system.dbConfig.oracle.enginePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.maxIdleConns')">
              <el-input-number
                v-model="config.oracle['max-idle-conns']"
                :min="1"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.maxOpenConns')">
              <el-input-number
                v-model="config.oracle['max-open-conns']"
                :min="1"
              />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.logZap')">
              <el-switch v-model="config.oracle['log-zap']" />
            </el-form-item>
            <el-form-item :label="$t('system.dbConfig.oracle.logMode')">
              <el-select v-model="config.oracle['log-mode']">
                <el-option
                  :value="'off'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.off')"
                />
                <el-option
                  :value="'fatal'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.fatal')"
                />
                <el-option
                  :value="'error'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.error')"
                />
                <el-option
                  :value="'warn'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.warn')"
                />
                <el-option
                  :value="'info'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.info')"
                />
                <el-option
                  :value="'debug'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.debug')"
                />
                <el-option
                  :value="'trace'"
                  :label="$t('system.dbConfig.mysql.logModeOptions.trace')"
                />
              </el-select>
            </el-form-item>
          </template>
        </el-tab-pane>
        <el-tab-pane
          :label="$t('system.ossConfig.title')"
          name="10"
          class="mt-3.5"
        >
          <template v-if="config.system['oss-type'] === 'local'">
            <h2>{{ $t('system.ossConfig.local.title') }}</h2>
            <el-form-item :label="$t('system.ossConfig.local.path')">
              <el-input
                v-model.trim="config.local.path"
                :placeholder="$t('system.ossConfig.local.pathPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.local.storePath')">
              <el-input
                v-model.trim="config.local['store-path']"
                :placeholder="$t('system.ossConfig.local.storePathPlaceholder')"
              />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'qiniu'">
            <h2>{{ $t('system.ossConfig.qiniu.title') }}</h2>
            <el-form-item :label="$t('system.ossConfig.qiniu.zone')">
              <el-input
                v-model.trim="config.qiniu.zone"
                :placeholder="$t('system.ossConfig.qiniu.zonePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.qiniu.bucket')">
              <el-input
                v-model.trim="config.qiniu.bucket"
                :placeholder="$t('system.ossConfig.qiniu.bucketPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.qiniu.imgPath')">
              <el-input
                v-model.trim="config.qiniu['img-path']"
                :placeholder="$t('system.ossConfig.qiniu.imgPathPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.qiniu.useHttps')">
              <el-switch v-model="config.qiniu['use-https']" />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.qiniu.accessKey')">
              <el-input
                v-model.trim="config.qiniu['access-key']"
                :placeholder="$t('system.ossConfig.qiniu.accessKeyPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.qiniu.secretKey')">
              <el-input
                v-model.trim="config.qiniu['secret-key']"
                :placeholder="$t('system.ossConfig.qiniu.secretKeyPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.qiniu.useCdnDomains')">
              <el-switch v-model="config.qiniu['use-cdn-domains']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'tencent-cos'">
            <h2>{{ $t('system.ossConfig.tencentCos.title') }}</h2>
            <el-form-item :label="$t('system.ossConfig.tencentCos.bucket')">
              <el-input
                v-model.trim="config['tencent-cos'].bucket"
                :placeholder="$t('system.ossConfig.tencentCos.bucketPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.tencentCos.region')">
              <el-input
                v-model.trim="config['tencent-cos'].region"
                :placeholder="$t('system.ossConfig.tencentCos.regionPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.tencentCos.secretId')">
              <el-input
                v-model.trim="config['tencent-cos']['secret-id']"
                :placeholder="$t('system.ossConfig.tencentCos.secretIdPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.tencentCos.secretKey')">
              <el-input
                v-model.trim="config['tencent-cos']['secret-key']"
                :placeholder="$t('system.ossConfig.tencentCos.secretKeyPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.tencentCos.pathPrefix')">
              <el-input
                v-model.trim="config['tencent-cos']['path-prefix']"
                :placeholder="$t('system.ossConfig.tencentCos.pathPrefixPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.tencentCos.baseUrl')">
              <el-input
                v-model.trim="config['tencent-cos']['base-url']"
                :placeholder="$t('system.ossConfig.tencentCos.baseUrlPlaceholder')"
              />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'aliyun-oss'">
            <h2>{{ $t('system.ossConfig.aliyunOss.title') }}</h2>
            <el-form-item :label="$t('system.ossConfig.aliyunOss.endpoint')">
              <el-input
                v-model.trim="config['aliyun-oss'].endpoint"
                :placeholder="$t('system.ossConfig.aliyunOss.endpointPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.aliyunOss.accessKeyId')">
              <el-input
                v-model.trim="config['aliyun-oss']['access-key-id']"
                :placeholder="$t('system.ossConfig.aliyunOss.accessKeyIdPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.aliyunOss.accessKeySecret')">
              <el-input
                v-model.trim="config['aliyun-oss']['access-key-secret']"
                :placeholder="$t('system.ossConfig.aliyunOss.accessKeySecretPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.aliyunOss.bucketName')">
              <el-input
                v-model.trim="config['aliyun-oss']['bucket-name']"
                :placeholder="$t('system.ossConfig.aliyunOss.bucketNamePlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.aliyunOss.bucketUrl')">
              <el-input
                v-model.trim="config['aliyun-oss']['bucket-url']"
                :placeholder="$t('system.ossConfig.aliyunOss.bucketUrlPlaceholder')"
              />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'huawei-obs'">
            <h2>{{ $t('system.ossConfig.huaweiObs.title') }}</h2>
            <el-form-item :label="$t('system.ossConfig.huaweiObs.path')">
              <el-input
                v-model.trim="config['huawei-obs'].path"
                :placeholder="$t('system.ossConfig.huaweiObs.pathPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.huaweiObs.bucket')">
              <el-input
                v-model.trim="config['huawei-obs'].bucket"
                :placeholder="$t('system.ossConfig.huaweiObs.bucketPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.huaweiObs.endpoint')">
              <el-input
                v-model.trim="config['huawei-obs'].endpoint"
                :placeholder="$t('system.ossConfig.huaweiObs.endpointPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.huaweiObs.accessKey')">
              <el-input
                v-model.trim="config['huawei-obs']['access-key']"
                :placeholder="$t('system.ossConfig.huaweiObs.accessKeyPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.huaweiObs.secretKey')">
              <el-input
                v-model.trim="config['huawei-obs']['secret-key']"
                :placeholder="$t('system.ossConfig.huaweiObs.secretKeyPlaceholder')"
              />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'cloudflare-r2'">
            <h2>{{ $t('system.ossConfig.cloudflareR2.title') }}</h2>
            <el-form-item :label="$t('system.ossConfig.cloudflareR2.path')">
              <el-input
                v-model.trim="config['cloudflare-r2'].path"
                :placeholder="$t('system.ossConfig.cloudflareR2.pathPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.cloudflareR2.bucket')">
              <el-input
                v-model.trim="config['cloudflare-r2'].bucket"
                :placeholder="$t('system.ossConfig.cloudflareR2.bucketPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.cloudflareR2.baseUrl')">
              <el-input
                v-model.trim="config['cloudflare-r2']['base-url']"
                :placeholder="$t('system.ossConfig.cloudflareR2.baseUrlPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.cloudflareR2.accountId')">
              <el-input
                v-model.trim="config['cloudflare-r2']['account-id']"
                :placeholder="$t('system.ossConfig.cloudflareR2.accountIdPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.cloudflareR2.accessKeyId')">
              <el-input
                v-model.trim="config['cloudflare-r2']['access-key-id']"
                :placeholder="$t('system.ossConfig.cloudflareR2.accessKeyIdPlaceholder')"
              />
            </el-form-item>
            <el-form-item :label="$t('system.ossConfig.cloudflareR2.secretAccessKey')">
              <el-input
                v-model.trim="config['cloudflare-r2']['secret-access-key']"
                :placeholder="$t('system.ossConfig.cloudflareR2.secretAccessKeyPlaceholder')"
              />
            </el-form-item>
          </template>
        </el-tab-pane>
        <el-tab-pane
          :label="$t('system.excelConfig.title')"
          name="11"
          class="mt-3.5"
        >
          <el-form-item :label="$t('system.excelConfig.targetDir')">
            <el-input
              v-model.trim="config.excel.dir"
              :placeholder="$t('system.excelConfig.targetDirPlaceholder')"
            />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
          :label="$t('system.autocodeConfig.title')"
          name="12"
          class="mt-3.5"
        >
          <el-form-item :label="$t('system.autocodeConfig.autoRestart')">
            <el-switch v-model="config.autocode['transfer-restart']" />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.root')">
            <el-input
              v-model="config.autocode.root"
              disabled
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.server.root')">
            <el-input
              v-model.trim="config.autocode['server']"
              :placeholder="$t('system.autocodeConfig.server.rootPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.server.api')">
            <el-input
              v-model.trim="config.autocode['server-api']"
              :placeholder="$t('system.autocodeConfig.server.apiPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.server.initialize')">
            <el-input
              v-model.trim="config.autocode['server-initialize']"
              :placeholder="$t('system.autocodeConfig.server.initializePlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.server.model')">
            <el-input
              v-model.trim="config.autocode['server-model']"
              :placeholder="$t('system.autocodeConfig.server.modelPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.server.request')">
            <el-input
              v-model.trim="config.autocode['server-request']"
              :placeholder="$t('system.autocodeConfig.server.requestPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.server.router')">
            <el-input
              v-model.trim="config.autocode['server-router']"
              :placeholder="$t('system.autocodeConfig.server.routerPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.server.service')">
            <el-input
              v-model.trim="config.autocode['server-service']"
              :placeholder="$t('system.autocodeConfig.server.servicePlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.web.root')">
            <el-input
              v-model.trim="config.autocode.web"
              :placeholder="$t('system.autocodeConfig.web.rootPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.web.api')">
            <el-input
              v-model.trim="config.autocode['web-api']"
              :placeholder="$t('system.autocodeConfig.web.apiPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.web.form')">
            <el-input
              v-model.trim="config.autocode['web-form']"
              :placeholder="$t('system.autocodeConfig.web.formPlaceholder')"
            />
          </el-form-item>
          <el-form-item :label="$t('system.autocodeConfig.web.table')">
            <el-input
              v-model.trim="config.autocode['web-table']"
              :placeholder="$t('system.autocodeConfig.web.tablePlaceholder')"
            />
          </el-form-item>
        </el-tab-pane>
      </el-tabs>
    </el-form>
    <div class="mt-4">
      <el-button
        type="primary"
        @click="update"
      >
        {{ $t('system.buttons.update') }}
      </el-button>
      <el-button
        type="primary"
        @click="reload"
      >
        {{ $t('system.buttons.restart') }}
      </el-button>
    </div>
  </div>
</template>

<script setup>
import { getSystemConfig, reloadSystem, setSystemConfig } from '@/api/system'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Minus, Plus } from '@element-plus/icons-vue'
import { CreateUUID } from '@/utils/format'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

defineOptions({
  name: 'Config'
})

const activeNames = ref('1')
const config = ref({
  system: {
    'iplimit-count': 0,
    'iplimit-time': 0
  },
  jwt: {},
  mysql: {},
  mssql: {},
  sqlite: {},
  pgsql: {},
  oracle: {},
  excel: {},
  autocode: {},
  redis: {},
  mongo: {
    coll: '',
    options: '',
    database: '',
    username: '',
    password: '',
    'min-pool-size': '',
    'max-pool-size': '',
    'socket-timeout-ms': '',
    'connect-timeout-ms': '',
    'is-zap': false,
    hosts: [
      {
        host: '',
        port: ''
      }
    ]
  },
  qiniu: {},
  'tencent-cos': {},
  'aliyun-oss': {},
  'hua-wei-obs': {},
  'cloudflare-r2': {},
  captcha: {},
  zap: {},
  local: {},
  email: {},
  timer: {
    detail: {}
  }
})

const initForm = async() => {
  const res = await getSystemConfig()
  if (res.code === 0) {
    config.value = res.data.config
  }
}
initForm()
const reload = () => {
  ElMessageBox.confirm(t('system.restartWarning.content'), t('system.restartWarning.title'), {
    confirmButtonText: t('system.restartWarning.confirm'),
    cancelButtonText: t('system.restartWarning.cancel'),
    type: 'warning'
  })
    .then(async() => {
      const res = await reloadSystem()
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('system.messages.restartSuccess')
        })
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: t('system.messages.restartCancel')
      })
    })
}

const update = async() => {
  const res = await setSystemConfig({ config: config.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: t('system.messages.updateSuccess')
    })
    await initForm()
  }
}


const getUUID = () => {
  config.value.jwt['signing-key'] = CreateUUID()
}

const addNode = () => {
  config.value.mongo.hosts.push({
    host: '',
    port: ''
  })
}

const removeNode = (index) => {
  config.value.mongo.hosts.splice(index, 1)
}
</script>

<style lang="scss" scoped>
  .system {
    @apply bg-white p-9 rounded dark:bg-slate-900;
    h2 {
      @apply p-2.5 my-2.5 text-lg shadow;
    }
  }
</style>
