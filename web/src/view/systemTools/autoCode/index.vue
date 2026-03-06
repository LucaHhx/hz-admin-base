<template>
  <div>
    <warning-bar
      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
      :title="$t('autoCode.warning')"
    />
    <div
      v-if="!isAdd"
      class="hab-search-box"
    >
      <!-- <div class="relative">
        <el-input
          v-model="prompt"
          type="textarea"
          :rows="5"
          :maxlength="2000"
          :placeholder="$t('autoCode.aiPlaceholder')"
          resize="none"
          @focus="handleFocus"
          @blur="handleBlur"
        />

        <div class="flex absolute right-28 bottom-2">
          <el-tooltip effect="light">
            <template #content>
              <div>
                {{ $t('autoCode.aiTip') }}
              </div>
            </template>
            <el-button
              :disabled="form.onlyTemplate"
              type="primary"
              @click="eyeFunc()"
            >
              <el-icon size="18">
                <ai-hab />
              </el-icon>
              {{ $t('autoCode.recognize') }}
            </el-button>
          </el-tooltip>
        </div>

        <div class="flex absolute right-2 bottom-2">
          <el-tooltip effect="light">
            <template #content>
              <div>
                {{ $t('autoCode.aiTip') }}
              </div>
            </template>
            <el-button
              :disabled="form.onlyTemplate"
              type="primary"
              @click="llmAutoFunc()"
            >
              <el-icon size="18">
                <ai-hab />
              </el-icon>
              {{ $t('autoCode.generate') }}
            </el-button>
          </el-tooltip>
        </div>
      </div> -->
    </div>
    <!-- 从数据库直接获取字段 -->
    <div
      v-if="!isAdd"
      class="hab-search-box"
    >
      <div class="text-lg mb-2 text-gray-600">
        {{ $t('autoCode.fromDb') }}
      </div>
      <el-form
        ref="getTableForm"
        :inline="true"
        :model="dbform"
        label-width="120px"
      >
        <el-row class="w-full">
          <el-col :span="6">
            <el-form-item
              label="业务库"
              prop="selectDBtype"
              class="w-full"
            >
              <template #label>
                <el-tooltip
                  :content="$t('autoCode.dbTip')"
                  placement="bottom"
                  effect="light"
                >
                  <div>
                    {{ $t('autoCode.businessDb') }} <el-icon><QuestionFilled /></el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-select
                v-model="dbform.businessDB"
                clearable
                :placeholder="$t('autoCode.selectDb')"
                class="w-full"
                @change="getDbFunc"
              >
                <el-option
                  v-for="item in dbList"
                  :key="item.aliasName"
                  :value="item.aliasName"
                  :label="item.aliasName"
                  :disabled="item.disable"
                >
                  <div>
                    <span>{{ item.aliasName }}</span>
                    <span
                      style="float: right; color: #8492a6; font-size: 13px"
                    >{{ item.dbName }}</span>
                  </div>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              :label="$t('autoCode.dbName')"
              prop="structName"
              class="w-full"
            >
              <el-select
                v-model="dbform.dbName"
                clearable
                filterable
                :placeholder="$t('autoCode.selectDbName')"
                class="w-full"
                @change="getTableFunc"
              >
                <el-option
                  v-for="item in dbOptions"
                  :key="item.database"
                  :label="item.database"
                  :value="item.database"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              :label="$t('autoCode.tableName')"
              prop="structName"
              class="w-full"
            >
              <el-select
                v-model="dbform.tableName"
                :disabled="!dbform.dbName"
                class="w-full"
                filterable
                :placeholder="$t('autoCode.selectTable')"
                @change="getColumnFunc"
              >
                <el-option
                  v-for="item in tableOptions"
                  :key="item.tableName"
                  :label="item.tableName"
                  :value="item.tableName"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item class="w-full">
              <div class="flex justify-end w-full">
                <el-button
                  type="primary"
                  @click="getColumnFunc"
                >
                  {{ $t('autoCode.useTable') }}
                </el-button>
              </div>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
    <div class="hab-search-box">
      <!-- 初始版本自动化代码工具 -->
      <div class="text-lg mb-2 text-gray-600">
        {{ $t('autoCode.autoStruct') }}
      </div>
      <el-form
        ref="autoCodeForm"
        :disabled="isAdd"
        :rules="rules"
        :model="form"
        label-width="120px"
        :inline="true"
      >
        <el-row class="w-full">
          <el-col :span="6">
            <el-form-item
              :label="$t('autoCode.structName')"
              prop="structName"
              class="w-full"
            >
              <div class="flex gap-2 w-full">
                <el-input
                  v-model="form.structName"
                  :placeholder="$t('autoCode.structNamePlaceholder')"
                />
                <el-button
                  :disabled="form.onlyTemplate"
                  type="primary"
                  @click="llmAutoFunc(true)"
                >
                  <el-icon size="18">
                    <ai-hab />
                  </el-icon>
                  {{ $t('autoCode.generate') }}
                </el-button>
              </div>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              label="TableName"
              class="w-full"
            >
              <template #label>
                <el-tooltip
                  :content="$t('autoCode.structAbbrTip')"
                  placement="bottom"
                  effect="light"
                >
                  <div>
                    {{ $t('autoCode.structAbbr') }} <el-icon><QuestionFilled /></el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-input
                v-model="form.abbreviation"
                :placeholder="$t('autoCode.structAbbrPlaceholder')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              :label="$t('autoCode.description')"
              prop="description"
              class="w-full"
            >
              <el-input
                v-model="form.description"
                :placeholder="$t('autoCode.descriptionPlaceholder')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              :label="$t('autoCode.tableName')"
              prop="tableName"
              class="w-full"
            >
              <el-input
                v-model="form.tableName"
                :placeholder="$t('autoCode.tableNamePlaceholder')"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="w-full">
          <el-col :span="6">
            <el-form-item
              prop="packageName"
              class="w-full"
            >
              <template #label>
                <el-tooltip
                  :content="$t('autoCode.packageNameTip')"
                  placement="bottom"
                  effect="light"
                >
                  <div>
                    {{ $t('autoCode.packageName') }} <el-icon><QuestionFilled /></el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-input
                v-model="form.packageName"
                :placeholder="$t('autoCode.packageNamePlaceholder')"
                @blur="toLowerCaseFunc(form, 'packageName')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              label="选择模板"
              prop="package"
              class="w-full relative"
            >
              <el-select
                v-model="form.package"
                class="w-full pr-12"
              >
                <el-option
                  v-for="item in pkgs"
                  :key="item.ID"
                  :value="item.packageName"
                  :label="item.packageName"
                />
              </el-select>
              <span class="absolute right-0">
                <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getPkgs"
                >
                  <refresh />
                </el-icon>
                <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goPkgs"
                >
                  <document-add />
                </el-icon>
              </span>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              label="业务库"
              prop="businessDB"
              class="w-full"
            >
              <template #label>
                <el-tooltip
                  :content="$t('autoCode.dbTip')"
                  placement="bottom"
                  effect="light"
                >
                  <div>
                    {{ $t('autoCode.businessDb') }} <el-icon><QuestionFilled /></el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-select
                v-model="form.businessDB"
                :placeholder="$t('autoCode.selectDb')"
                class="w-full"
              >
                <el-option
                  v-for="item in dbList"
                  :key="item.aliasName"
                  :value="item.aliasName"
                  :label="item.aliasName"
                  :disabled="item.disable"
                >
                  <div>
                    <span>{{ item.aliasName }}</span>
                    <span
                      style="float: right; color: #8492a6; font-size: 13px"
                    >{{ item.dbName }}</span>
                  </div>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
    <div class="hab-search-box">
      <el-collapse class="no-border-collapse">
        <el-collapse-item>
          <template #title>
            <div class="text-lg text-gray-600 font-normal">
              {{ $t('autoCode.expertMode') }}
            </div>
          </template>
          <template #icon="{ isActive }">
            <span class="text-lg ml-auto mr-4 font-normal">
              {{ isActive ? $t('autoCode.index.collapse') : $t('autoCode.index.expand') }}
            </span>
          </template>
          <div class="p-4">
            <!-- 基础设置组 -->
            <div class="border-b border-gray-200 last:border-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">
                {{ $t('autoCode.basicSettings') }}
              </h3>
              <el-row :gutter="20">
                <el-col :span="3">
                  <el-tooltip
                    :content="$t('autoCode.habModelTip')"
                    placement="top"
                    effect="light"
                  >
                    <el-form-item :label="$t('autoCode.index.useDefaultStructure')">
                      <el-checkbox
                        v-model="form.habModel"
                        @change="useHab"
                      />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                    :content="$t('autoCode.createBtnAuthTip')"
                    placement="top"
                    effect="light"
                  >
                    <el-form-item :label="$t('autoCode.index.createButtonAuth')">
                      <el-checkbox
                        v-model="form.autoCreateBtnAuth"
                        :disabled="!form.generateWeb"
                      />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                    :content="$t('autoCode.createBtnAuthTip')"
                    placement="top"
                    effect="light"
                  >
                    <el-form-item :label="$t('autoCode.index.generateFrontend')">
                      <el-checkbox v-model="form.generateWeb" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                    :content="$t('autoCode.createBtnAuthTip')"
                    placement="top"
                    effect="light"
                  >
                    <el-form-item :label="$t('autoCode.index.generateBackend')">
                      <el-checkbox
                        v-model="form.generateServer"
                        disabled
                      />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
              </el-row>
            </div>

            <!-- 自动化设置组 -->
            <div class="border-b border-gray-200 last:border-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">
                {{ $t('autoCode.autoSettings') }}
              </h3>
              <el-row :gutter="20">
                <el-col :span="3">
                  <el-tooltip
                    :content="$t('autoCode.createApiTip')"
                    placement="top"
                    effect="light"
                  >
                    <el-form-item :label="$t('autoCode.index.autoCreateAPI')">
                      <el-checkbox
                        v-model="form.autoCreateApiToSql"
                        :disabled="!form.generateServer"
                      />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                    :content="$t('autoCode.createMenuTip')"
                    placement="top"
                    effect="light"
                  >
                    <el-form-item :label="$t('autoCode.index.autoCreateMenu')">
                      <el-checkbox
                        v-model="form.autoCreateMenuToSql"
                        :disabled="!form.generateWeb"
                      />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                    :content="$t('autoCode.syncTableTip')"
                    placement="top"
                    effect="light"
                  >
                    <el-form-item :label="$t('autoCode.index.syncTableStructure')">
                      <el-checkbox
                        v-model="form.autoMigrate"
                        :disabled="!form.generateServer"
                      />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
              </el-row>
            </div>

            <!-- 高级设置组 -->
            <div class="border-b border-gray-200 last:border-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">
                {{ $t('autoCode.advancedSettings') }}
              </h3>
              <el-row :gutter="20">
                <el-col :span="3">
                  <el-tooltip
                    :content="$t('autoCode.createResourceTip')"
                    placement="top"
                    effect="light"
                  >
                    <el-form-item label="创建资源标识">
                      <el-checkbox v-model="form.autoCreateResource" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                    :content="$t('autoCode.useTemplateTip')"
                    placement="top"
                    effect="light"
                  >
                    <el-form-item label="基础模板">
                      <el-checkbox v-model="form.onlyTemplate" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
              </el-row>
            </div>

            <!-- 树形结构设置 -->
            <div class="last:pb-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">
                {{ $t('autoCode.treeStructSettings') }}
              </h3>
              <el-row
                :gutter="20"
                align="middle"
              >
                <el-col :span="24">
                  <el-form-item label="树型结构">
                    <div class="flex items-center gap-4">
                      <el-tooltip
                        :content="$t('autoCode.treeTip')"
                        placement="top"
                        effect="light"
                      >
                        <el-checkbox v-model="form.isTree" />
                      </el-tooltip>
                      <el-input
                        v-model="form.treeJson"
                        :disabled="!form.isTree"
                        :placeholder="$t('autoCode.treeJsonPlaceholder')"
                        class="flex-1"
                      />
                    </div>
                  </el-form-item>
                </el-col>
              </el-row>
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <!-- 组件列表 -->
    <div class="hab-table-box">
      <div class="hab-btn-list">
        <el-button
          type="primary"
          :disabled="form.onlyTemplate"
          @click="editAndAddField()"
        >
          {{ $t('autoCode.addField') }}
        </el-button>
      </div>
      <div class="draggable">
        <el-table
          :data="form.fields"
          row-key="fieldName"
          table-layout="auto"
          max-height="500"
        >
          <el-table-column
            v-if="!isAdd"
            fixed="left"
            align="left"
            type="index"
            :width="null"
          >
            <template #default>
              <el-icon class="cursor-grab drag-column">
                <MoreFilled />
              </el-icon>
            </template>
          </el-table-column>
          <el-table-column
            fixed="left"
            align="left"
            type="index"
            :label="$t('autoCode.index.sequence')"
            width="60"
          />
          <el-table-column
            fixed="left"
            align="left"
            type="index"
            :label="$t('autoCode.index.primaryKey')"
            :width="null"
          >
            <template #default="{ row }">
              <el-checkbox
                v-model="row.primaryKey"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            fixed="left"
            align="left"
            prop="fieldName"
            :label="$t('autoCode.index.fieldName')"
            width="160"
          >
            <template #default="{ row }">
              <el-input
                v-model="row.fieldName"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldDesc"
            :label="$t('autoCode.index.chineseName')"
            width="160"
          >
            <template #default="{ row }">
              <el-input
                v-model="row.fieldDesc"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="defaultValue"
            :label="$t('autoCode.index.defaultValue')"
            width="160"
          >
            <template #default="{ row }">
              <el-input
                v-model="row.defaultValue"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="require"
            :label="$t('autoCode.index.required')"
          >
            <template #default="{ row }">
              <el-checkbox
                v-model="row.require"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="sort"
            :label="$t('autoCode.index.sort')"
          >
            <template #default="{ row }">
              <el-checkbox
                v-model="row.sort"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="form"
            width="100"
            :label="$t('autoCode.index.createEdit')"
          >
            <template #default="{ row }">
              <el-checkbox
                v-model="row.form"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="table"
            :label="$t('autoCode.index.table')"
          >
            <template #default="{ row }">
              <el-checkbox
                v-model="row.table"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="desc"
            :label="$t('autoCode.index.detail')"
          >
            <template #default="{ row }">
              <el-checkbox
                v-model="row.desc"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            v-if="!isAdd"
            align="left"
            prop="excel"
            width="100"
            :label="$t('autoCode.index.importExport')"
          >
            <template #default="{ row }">
              <el-checkbox v-model="row.excel" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldJson"
            width="160px"
            :label="$t('autoCode.index.fieldJson')"
          >
            <template #default="{ row }">
              <el-input
                v-model="row.fieldJson"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldType"
            :label="$t('autoCode.index.fieldType')"
            width="160"
          >
            <template #default="{ row }">
              <el-select
                v-model="row.fieldType"
                style="width: 100%"
                :placeholder="$t('autoCode.index.selectFieldType')"
                :disabled="row.disabled"
                clearable
              >
                <el-option
                  v-for="item in typeOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldIndexType"
            :label="$t('autoCode.index.indexType')"
            width="160"
          >
            <template #default="{ row }">
              <el-select
                v-model="row.fieldIndexType"
                style="width: 100%"
                placeholder="请选择字段索引类型"
                :disabled="row.disabled"
                clearable
              >
                <el-option
                  v-for="item in typeIndexOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="dataTypeLong"
            :label="$t('autoCode.index.dbFieldLength')"
            width="160"
          >
            <template #default="{ row }">
              <el-input
                v-model="row.dataTypeLong"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="columnName"
            :label="$t('autoCode.index.dbField')"
            width="160"
          >
            <template #default="{ row }">
              <el-input
                v-model="row.columnName"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="comment"
            :label="$t('autoCode.index.dbFieldDesc')"
            width="160"
          >
            <template #default="{ row }">
              <el-input
                v-model="row.comment"
                :disabled="row.disabled"
              />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldSearchType"
            :label="$t('autoCode.index.searchCondition')"
            width="130"
          >
            <template #default="{ row }">
              <el-select
                v-model="row.fieldSearchType"
                style="width: 100%"
                :placeholder="$t('autoCode.index.selectSearchCondition')"
                clearable
                :disabled="row.fieldType == 'json' || row.disabled"
              >
                <el-option
                  v-for="item in typeSearchOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                  :disabled="canSelect(row.fieldType,item.value)"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="$t('autoCode.index.operations')"
            width="300"
            fixed="right"
          >
            <template #default="scope">
              <el-button
                v-if="!scope.row.disabled"
                type="primary"
                link
                icon="edit"
                @click="editAndAddField(scope.row)"
              >
                {{ $t('autoCode.advancedEdit') }}
              </el-button>
              <el-button
                v-if="!scope.row.disabled"
                type="primary"
                link
                icon="delete"
                @click="deleteField(scope.$index)"
              >
                {{ $t('autoCode.delete') }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!-- 组件列表 -->
      <div class="hab-btn-list justify-end mt-4">
        <el-button
          type="primary"
          :disabled="isAdd"
          @click="exportJson()"
        >
          {{ $t('autoCode.exportJson') }}
        </el-button>
        <el-upload
          class="flex items-center"
          :before-upload="importJson"
          :show-file-list="false"
          accept=".json"
        >
          <el-button
            type="primary"
            class="mx-2"
            :disabled="isAdd"
          >
            {{ $t('autoCode.importJson') }}
          </el-button>
        </el-upload>
        <el-button
          type="primary"
          :disabled="isAdd"
          @click="clearCatch()"
        >
          {{ $t('autoCode.clearCatch') }}
        </el-button>
        <el-button
          type="primary"
          :disabled="isAdd"
          @click="catchData()"
        >
          {{ $t('autoCode.catchData') }}
        </el-button>
        <el-button
          type="primary"
          :disabled="isAdd"
          @click="enterForm(false)"
        >
          {{ $t('autoCode.generateCode') }}
        </el-button>
        <el-button
          type="primary"
          @click="enterForm(true)"
        >
          {{ isAdd ? $t('autoCode.viewCode') : $t('autoCode.previewCode') }}
        </el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-drawer
      v-model="dialogFlag"
      size="70%"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ $t('autoCode.componentContent') }}</span>
          <div>
            <el-button @click="closeDialog">
              {{ $t('autoCode.cancel') }}
            </el-button>
            <el-button
              type="primary"
              @click="enterDialog"
            >
              {{ $t('autoCode.confirm') }}
            </el-button>
          </div>
        </div>
      </template>

      <FieldDialog
        v-if="dialogFlag"
        ref="fieldDialogNode"
        :dialog-middle="dialogMiddle"
        :type-options="typeOptions"
        :type-search-options="typeSearchOptions"
        :type-index-options="typeIndexOptions"
      />
    </el-drawer>

    <el-drawer
      v-model="previewFlag"
      size="80%"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ $t('autoCode.operationBar') }}</span>
          <div>
            <el-button
              type="primary"
              @click="selectText"
            >
              {{ $t('autoCode.selectAll') }}
            </el-button>
            <el-button
              type="primary"
              @click="copy"
            >
              {{ $t('autoCode.copy') }}
            </el-button>
          </div>
        </div>
      </template>
      <PreviewCodeDialog
        v-if="previewFlag"
        ref="previewNode"
        :is-add="isAdd"
        :preview-code="preViewCode"
      />
    </el-drawer>
  </div>
</template>

<script setup>
import FieldDialog from '@/view/systemTools/autoCode/component/fieldDialog.vue'
import PreviewCodeDialog from '@/view/systemTools/autoCode/component/previewCodeDialog.vue'
import {
  toUpperCase,
  toHump,
  toSQLLine,
  toLowerCase
} from '@/utils/stringFun'
import {
  createTemp,
  getDB,
  getTable,
  getColumn,
  preview,
  getMeta,
  getPackageApi,
  llmAuto, eye
} from '@/api/autoCode'
import { getDict } from '@/utils/dictionary'
import { ref, watch, toRaw, onMounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'
import Sortable from 'sortablejs'

const handleFocus = () => {
  document.addEventListener('keydown', handleKeydown)
  document.addEventListener('paste', handlePaste)
}

const handleBlur = () => {
  document.removeEventListener('keydown', handleKeydown)
  document.removeEventListener('paste', handlePaste)
}

const handleKeydown = (event) => {
  if ((event.ctrlKey || event.metaKey) && event.key === 'Enter') {
    llmAutoFunc()
  }
}

const handlePaste = (event) => {
  const items = event.clipboardData.items
  for (let i = 0; i < items.length; i++) {
    if (items[i].type.indexOf('image') !== -1) {
      const file = items[i].getAsFile()
      const reader = new FileReader()
      reader.onload = async(e) => {
        const base64String = e.target.result
        const res = await eye({ picture: base64String, command: 'eye' })
        if (res.code === 0) {
          prompt.value = res.data
          llmAutoFunc()
        }
      }
      reader.readAsDataURL(file)
    }
  }
}

const getOnlyNumber = () => {
  let randomNumber = ''
  while (randomNumber.length < 16) {
    randomNumber += Math.random().toString(16).substring(2)
  }
  return randomNumber.substring(0, 16)
}

const prompt = ref('')

const eyeFunc = async() => {
  const input = document.createElement('input')
  input.type = 'file'
  input.accept = 'image/*'

  input.onchange = (event) => {
    const file = event.target.files[0]
    if (file) {
      const reader = new FileReader()
      reader.onload = async(e) => {
        const base64String = e.target.result

        const res = await eye({ picture: base64String, command: 'eye' })
        if (res.code === 0) {
          prompt.value = res.data
          llmAutoFunc()
        }
      }
      reader.readAsDataURL(file)
    }
  }

  input.click()
}

const llmAutoFunc = async(flag) => {
  if (flag && !form.value.structName) {
    ElMessage.error('请输入结构体名称')
    return
  }
  if (!flag && !prompt.value) {
    ElMessage.error('请输入描述')
    return
  }

  if (form.value.fields.length > 0) {
    const res = await ElMessageBox.confirm(
      'AI生成会清空当前数据，是否继续?',
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    if (res !== 'confirm') {
      return
    }
  }

  const res = await llmAuto({
    prompt: flag ? '结构体名称为' + form.value.structName : prompt.value
  })
  if (res.code === 0) {
    form.value.fields = []
    const json = JSON.parse(res.data)

    json.fields?.forEach((item) => {
      item.fieldName = toUpperCase(item.fieldName)
    })

    for (const key in json) {
      form.value[key] = json[key]
    }

    form.value.generateServer = true
    form.value.generateWeb = true
  }
}

const isAdd = ref(false)

// 行拖拽
const rowDrop = () => {
  // 要拖拽元素的父容器
  const tbody = document.querySelector(
    '.draggable .el-table__body-wrapper tbody'
  )
  Sortable.create(tbody, {
    //  可被拖拽的子元素
    draggable: '.draggable .el-table__row',
    handle: '.drag-column',
    onEnd: async({ newIndex, oldIndex }) => {
      await nextTick()
      const currRow = form.value.fields.splice(oldIndex, 1)[0]
      form.value.fields.splice(newIndex, 0, currRow)
    }
  })
}

onMounted(() => {
  rowDrop()
})

defineOptions({
  name: 'AutoCode'
})
const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']

const dataModelList = ['created_by', 'updated_by', 'deleted_by']

const typeOptions = ref([
  {
    label: '字符串',
    value: 'string'
  },
  {
    label: '富文本',
    value: 'richtext'
  },
  {
    label: '布尔值',
    value: 'bool'
  },
  {
    label: '枚举',
    value: 'enum'
  },
  {
    label: '整型',
    value: 'int32'
  },
  {
    label: '大整数',
    value: 'int64'
  },
  {
    label: '浮点型',
    value: 'float64'
  },
  {
    label: '整型',
    value: 'int'
  },
  {
    label: '时间',
    value: 'datetime'
  },
  {
    label: '日期',
    value: 'date'
  },
  {
    label: 'JSON',
    value: 'json'
  },
  {
    label: '数组',
    value: 'array'
  },
  {
    label: '二进制',
    value: 'binary'
  },

])

const dbToGoTypeMap = {

  tinyint: 'bool',
  smallint: 'enum',
  mediumint: 'int32',
  int: 'int32',
  bigint: 'int64',
  timestamp: 'int64',

  double: 'float64',
  float: 'float64',

  datetime: 'datetime',
  date: 'date',

  varchar: 'string',
  char: 'string',

  tinytext: 'richtext',
  text: 'richtext',
  mediumtext: 'richtext',
  longtext: 'richtext',

  json: 'json',

  binary: 'binary',
  tinyblob: 'binary',
  blob: 'binary',
  mediumblob: 'binary',
  longblob: 'binary',
}

const searchTypeMap = {
  'string': '=',
  'richtext': 'like',
  'bool': '=',
  'enum': 'IN',
  'int32': 'BETWEEN',
  'int64': 'BETWEEN',
  'float64': 'BETWEEN',
  'datetime': 'BETWEEN',
  'date': 'BETWEEN',
}

const sortTypeList = [
  'string',
  'bool',
  'enum',
  'int32',
  'int64',
  'float64',
  'datetime',
  'date',
]

const typeSearchOptions = ref([
  {
    label: '=',
    value: '='
  },
  {
    label: '<>',
    value: '<>'
  },
  {
    label: '>',
    value: '>'
  },
  {
    label: '<',
    value: '<'
  },
  {
    label: 'LIKE',
    value: 'LIKE'
  },
  {
    label: 'BETWEEN',
    value: 'BETWEEN'
  },
  {
    label: 'NOT BETWEEN',
    value: 'NOT BETWEEN'
  },
  {
    label: 'IN',
    value: 'IN'
  },
  {
    label: 'NOT IN',
    value: 'NOT IN'
  }
])

const typeIndexOptions = ref([
  {
    label: 'index',
    value: 'index'
  },
  {
    label: 'uniqueIndex',
    value: 'uniqueIndex'
  }
])

const fieldTemplate = {
  fieldName: '',
  fieldDesc: '',
  fieldType: '',
  dataType: '',
  fieldJson: '',
  columnName: '',
  dataTypeLong: '',
  comment: '',
  defaultValue: '',
  require: false,
  sort: false,
  form: true,
  desc: true,
  table: true,
  excel: false,
  errorText: '',
  primaryKey: false,
  clearable: true,
  fieldSearchType: '',
  fieldIndexType: '',
  dictType: '',
  dataSource: {
    dbName: '',
    association: 1,
    table: '',
    label: '',
    value: '',
    hasDeletedAt: false
  }
}
const route = useRoute()
const router = useRouter()
const preViewCode = ref({})
const dbform = ref({
  businessDB: '',
  dbName: '',
  tableName: ''
})
const tableOptions = ref([])
const addFlag = ref('')

const form = ref({
  structName: '',
  tableName: '',
  packageName: '',
  package: 'business',
  abbreviation: '',
  description: '',
  businessDB: '',
  autoCreateApiToSql: true,
  autoCreateMenuToSql: true,
  autoCreateBtnAuth: true,
  autoMigrate: true,
  habModel: true,
  autoCreateResource: true,
  onlyTemplate: false,
  isTree: false,
  generateWeb: true,
  generateServer: true,
  treeJson: '',
  fields: []
})
const rules = ref({
  structName: [
    { required: true, message: '请输入结构体名称', trigger: 'blur' }
  ],
  abbreviation: [
    { required: true, message: '请输入结构体简称', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入结构体描述', trigger: 'blur' }
  ],
  packageName: [
    {
      required: true,
      message: '文件名称：sysXxxxXxxx',
      trigger: 'blur'
    }
  ],
  package: [{ required: true, message: '请选择package', trigger: 'blur' }]
})
const dialogMiddle = ref({})
const bk = ref({})
const dialogFlag = ref(false)
const previewFlag = ref(false)

const useHab = (e) => {
  if (e && form.value.fields.length) {
    ElMessageBox.confirm(
      '如果您开启默认结构，会自动添加ID,CreatedAt,UpdatedAt,DeletedAt字段，此行为将自动清除您目前在下方创建的重名字段，是否继续？',
      '注意',
      {
        confirmButtonText: '继续',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
      .then(() => {
        form.value.fields = form.value.fields.filter(
          (item) =>
            !gormModelList.some((gormfd) => gormfd === item.columnName)
        )
      })
      .catch(() => {
        form.value.habModel = false
      })
  }
}

const toLowerCaseFunc = (form, key) => {
  form[key] = toLowerCase(form[key])
}
const previewNode = ref(null)
const selectText = () => {
  previewNode.value.selectText()
}
const copy = () => {
  previewNode.value.copy()
}
const editAndAddField = (item) => {
  dialogFlag.value = true
  if (item) {
    addFlag.value = 'edit'
    if (!item.dataSource) {
      item.dataSource = {
        dbName: '',
        association: 1,
        table: '',
        label: '',
        value: '',
        hasDeletedAt: false
      }
    }
    bk.value = JSON.parse(JSON.stringify(item))
    dialogMiddle.value = item
  } else {
    addFlag.value = 'add'
    fieldTemplate.onlyNumber = getOnlyNumber()
    dialogMiddle.value = JSON.parse(JSON.stringify(fieldTemplate))
  }
}

const fieldDialogNode = ref(null)
const enterDialog = () => {
  fieldDialogNode.value.fieldDialogForm.validate((valid) => {
    if (valid) {
      dialogMiddle.value.fieldName = toUpperCase(dialogMiddle.value.fieldName)
      if (addFlag.value === 'add') {
        form.value.fields.push(dialogMiddle.value)
      }
      dialogFlag.value = false
    } else {
      return false
    }
  })
}
const closeDialog = () => {
  if (addFlag.value === 'edit') {
    dialogMiddle.value = bk.value
  }
  dialogFlag.value = false
}
const deleteField = (index) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    form.value.fields.splice(index, 1)
  })
}
const autoCodeForm = ref(null)
const enterForm = async(isPreview) => {
  if (form.value.isTree && !form.value.treeJson) {
    ElMessage({
      type: 'error',
      message: '请填写树型结构的前端展示json属性'
    })
    return false
  }
  if (!form.value.generateWeb && !form.value.generateServer) {
    ElMessage({
      type: 'error',
      message: '请至少选择一个生成项'
    })
    return false
  }
  if (!form.value.onlyTemplate) {
    if (form.value.fields.length <= 0) {
      ElMessage({
        type: 'error',
        message: '请填写至少一个field'
      })
      return false
    }

    if (
      !form.value.habModel &&
        form.value.fields.every((item) => !item.primaryKey)
    ) {
      ElMessage({
        type: 'error',
        message: '您至少需要创建一个主键才能保证自动化代码的可行性'
      })
      return false
    }

    if (
      form.value.fields.some(
        (item) => item.fieldName === form.value.structName
      )
    ) {
      ElMessage({
        type: 'error',
        message: '存在与结构体同名的字段'
      })
      return false
    }

    if (
      form.value.fields.some((item) => item.fieldJson === form.value.package)
    ) {
      ElMessage({
        type: 'error',
        message: '存在与模板同名的的字段JSON'
      })
      return false
    }

    if (form.value.fields.some((item) => !item.fieldType)) {
      ElMessage({
        type: 'error',
        message: '请填写所有字段类型后进行提交'
      })
      return false
    }

    if (form.value.package === form.value.abbreviation) {
      ElMessage({
        type: 'error',
        message: 'package和结构体简称不可同名'
      })
      return false
    }
  }

  autoCodeForm.value.validate(async(valid) => {
    if (valid) {
      for (const key in form.value) {
        if (typeof form.value[key] === 'string') {
          form.value[key] = form.value[key].trim()
        }
      }
      form.value.structName = toUpperCase(form.value.structName)
      form.value.tableName = form.value.tableName.replace(' ', '')
      if (!form.value.tableName) {
        form.value.tableName = toSQLLine(toLowerCase(form.value.structName))
      }
      if (form.value.structName === form.value.abbreviation) {
        ElMessage({
          type: 'error',
          message: 'structName和struct简称不能相同'
        })
        return false
      }
      form.value.humpPackageName = toSQLLine(form.value.packageName)

      form.value.fields?.forEach((item) => {
        item.fieldName = toUpperCase(item.fieldName)
        // if (item.fieldType === 'enum') {
        //   // 判断一下 item.dataTypeLong 按照,切割后的每个元素是否都使用 '' 包裹，如果没包 则修改为包裹起来的 然后再转为字符串赋值给 item.dataTypeLong
        //   // item.dataTypeLong = item.dataTypeLong.replace(/[[\]{}()]/g, '')
        //   const arr = item.dataTypeLong.split(',')
        //   arr.forEach((ele, index) => {
        //     if (ele.indexOf("'") === -1) {
        //       arr[index] = `'${ele}'`
        //     }
        //   })
        //   item.dataTypeLong = arr.join(',')
        // }
      })

      delete form.value.primaryField
      if (isPreview) {
        const res = await preview({
          ...form.value,
          isAdd: !!isAdd.value,
          fields: form.value.fields.filter((item) => !item.disabled)
        })
        if (res.code !== 0) {
          return
        }
        preViewCode.value = res.data.autoCode
        previewFlag.value = true
      } else {
        const res = await createTemp(form.value)
        if (res.code !== 0) {
          return
        }
        ElMessage({
          type: 'success',
          message: '自动化代码创建成功，自动移动成功'
        })
        clearCatch()
      }
    }
  })
}

const dbList = ref([])
const dbOptions = ref([])

const getDbFunc = async() => {
  dbform.value.dbName = ''
  dbform.value.tableName = ''
  const res = await getDB({ businessDB: dbform.value.businessDB })
  if (res.code === 0) {
    dbOptions.value = res.data.dbs
    dbList.value = res.data.dbList
  }
}
const getTableFunc = async() => {
  const res = await getTable({
    businessDB: dbform.value.businessDB,
    dbName: dbform.value.dbName
  })
  if (res.code === 0) {
    tableOptions.value = res.data.tables
  }
  dbform.value.tableName = ''
}

const getColumnFunc = async() => {
  const res = await getColumn(dbform.value)
  if (res.code === 0) {
    let dbtype = ''
    if (dbform.value.businessDB !== '') {
      const dbtmp = dbList.value.find(
        (item) => item.aliasName === dbform.value.businessDB
      )
      const dbraw = toRaw(dbtmp)
      dbtype = dbraw.dbtype
    }
    form.value.habModel = true
    const tbHump = toHump(dbform.value.tableName)
    form.value.structName = toUpperCase(tbHump)
    form.value.businessDB = dbform.value.businessDB
    form.value.tableName = dbform.value.tableName
    form.value.packageName = toLowerCase(tbHump)
    form.value.abbreviation = toLowerCase(tbHump)
    form.value.description = tbHump
    form.value.autoCreateApiToSql = true
    form.value.generateServer = true
    form.value.generateWeb = true
    form.value.fields = []
    console.log(res.data.columns)

    res.data.columns &&
        res.data.columns.forEach((item) => {
          if (needAppend(item)) {
            const fbHump = toHump(item.columnName)
            form.value.fields.push({
              onlyNumber: getOnlyNumber(),
              fieldName: toUpperCase(fbHump),
              fieldDesc: item.columnComment || fbHump + '字段',
              fieldType: dbToGoTypeMap[item.dataType],
              dataType: item.dataType,
              fieldJson: fbHump,
              primaryKey: item.primaryKey,
              dataTypeLong: dbToGoTypeMap[item.dataType] === 'string' ? item.dataTypeLong : null,
              columnName:
                dbtype === 'oracle'
                  ? item.columnName.toUpperCase()
                  : item.columnName,
              sort: sortTypeList.includes(dbToGoTypeMap[item.dataType]),
              comment: item.columnComment,
              require: false,
              errorText: '',
              clearable: true,
              fieldSearchType: searchTypeMap[dbToGoTypeMap[item.dataType]],
              fieldIndexType: '',
              dictType: '',
              form: true,
              table: true,
              excel: false,
              desc: true,
              dataSource: {
                dbName: '',
                association: 1,
                table: '',
                label: '',
                value: '',
                hasDeletedAt: false
              }
            })
          }
        })
  }
}

const needAppend = (item) => {
  let isAppend = true
  if (
    form.value.habModel &&
      gormModelList.some((gormfd) => gormfd === item.columnName)
  ) {
    isAppend = false
  }
  if (
    form.value.autoCreateResource &&
      dataModelList.some((datafd) => datafd === item.columnName)
  ) {
    isAppend = false
  }
  return isAppend
}

const getAutoCodeJson = async(id) => {
  const res = await getMeta({ id: Number(id) })
  if (res.code === 0) {
    const add = route.query.isAdd
    isAdd.value = add
    form.value = JSON.parse(res.data.meta)
    if (isAdd.value) {
      form.value.fields.forEach((item) => {
        item.disabled = true
      })
    }
  }
}

const pkgs = ref([])
const getPkgs = async() => {
  const res = await getPackageApi()
  if (res.code === 0) {
    pkgs.value = res.data.pkgs
  }
}

const goPkgs = () => {
  router.push({ name: 'autoPkg' })
}

const init = () => {
  getDbFunc()
  getPkgs()
  const id = route.params.id
  if (id) {
    getAutoCodeJson(id)
  }
}
init()

watch(
  () => route.params.id,
  () => {
    if (route.name === 'autoCodeEdit') {
      init()
    }
  }
)

watch(() => form.value.generateServer, () => {
  if (!form.value.generateServer) {
    form.value.autoCreateApiToSql = false
    form.value.autoMigrate = false
  }
})

watch(() => form.value.generateWeb, () => {
  if (!form.value.generateWeb) {
    form.value.autoCreateMenuToSql = false
    form.value.autoCreateBtnAuth = false
  }
})

const catchData = () => {
  window.sessionStorage.setItem('autoCode', JSON.stringify(form.value))
}

const getCatch = () => {
  const data = window.sessionStorage.getItem('autoCode')
  if (data) {
    form.value = JSON.parse(data)
  }
}

const clearCatch = async() => {
  form.value = {
    structName: '',
    tableName: '',
    packageName: '',
    package: '',
    abbreviation: '',
    description: '',
    businessDB: '',
    autoCreateApiToSql: true,
    autoCreateMenuToSql: true,
    autoCreateBtnAuth: false,
    autoMigrate: true,
    habModel: true,
    autoCreateResource: false,
    onlyTemplate: false,
    isTree: false,
    treeJson: '',
    fields: []
  }
  await nextTick()
  window.sessionStorage.removeItem('autoCode')
}

getCatch()

const exportJson = () => {
  const dataStr = JSON.stringify(form.value, null, 2)
  const blob = new Blob([dataStr], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = 'form_data.json'
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)
}

const importJson = (file) => {
  const reader = new FileReader()
  reader.onload = (e) => {
    try {
      form.value = JSON.parse(e.target.result)
      ElMessage.success('JSON 文件导入成功')
    } catch {
      ElMessage.error('无效的 JSON 文件')
    }
  }
  reader.readAsText(file)
  return false
}

watch(
  () => form.value.onlyTemplate,
  (val) => {
    if (val) {
      ElMessageBox.confirm(
        '使用基础模板将不会生成任何结构体和CURD,仅仅配置enter等属性方便自行开发非CURD逻辑',
        '注意',
        {
          confirmButtonText: '继续',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
        .then(() => {
          form.value.fields = []
        })
        .catch(() => {
          form.value.onlyTemplate = false
        })
    }
  }
)

const canSelect = (fieldType, item) => {
  if (fieldType === 'richtext') {
    return item !== 'LIKE'
  }

  if (fieldType !== 'string' && item === 'LIKE') {
    return true
  }

  const nonNumericTypes = ['int', 'time.Time', 'float64']
  if (!nonNumericTypes.includes(fieldType) && ['BETWEEN', 'NOT BETWEEN'].includes(item)) {
    return true
  }

  return false
}
</script>

<style>
.no-border-collapse{
  @apply border-none;
  .el-collapse-item__header{
    @apply border-none;
  }
  .el-collapse-item__wrap{
    @apply border-none;
  }
  .el-collapse-item__content{
    @apply pb-0;
  }
}
</style>
