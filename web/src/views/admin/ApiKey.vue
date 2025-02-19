<template>
  <div class="container list" v-loading="loading">

    <div class="handle-box">
      <el-button type="primary" :icon="Plus" @click="add">新增</el-button>
    </div>

    <el-row>
      <el-table :data="items" :row-key="row => row.id" table-layout="auto">
        <el-table-column prop="platform" label="所属平台"/>
        <el-table-column prop="value" label="KEY"/>

        <el-table-column label="创建时间">
          <template #default="scope">
            <span>{{ dateFormat(scope.row['created_at']) }}</span>
          </template>
        </el-table-column>

        <el-table-column label="最后使用时间">
          <template #default="scope">
            <span v-if="scope.row['last_used_at']">{{ scope.row['last_used_at'] }}</span>
            <el-tag v-else>未使用</el-tag>
          </template>
        </el-table-column>

        <el-table-column label="操作" width="180">
          <template #default="scope">
            <el-button size="small" type="primary" @click="edit(scope.row)">编辑</el-button>
            <el-popconfirm title="确定要删除当前记录吗?" @confirm="remove(scope.row)">
              <template #reference>
                <el-button size="small" type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-row>

    <el-dialog
        v-model="showDialog"
        :title="title"
        style="width: 90%; max-width: 600px;"
    >
      <el-form :model="item" label-width="120px" ref="formRef" :rules="rules">
        <el-form-item label="所属平台：" prop="platform">
          <el-select v-model="item.platform" placeholder="请选择平台">
            <el-option v-for="item in platforms" :value="item" :key="item">{{ item }}</el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="API KEY：" prop="value">
          <el-input v-model="item.value" autocomplete="off"/>
        </el-form-item>

      </el-form>

      <template #footer>
            <span class="dialog-footer">
              <el-button @click="showDialog = false">取消</el-button>
              <el-button type="primary" @click="save">提交</el-button>
            </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {reactive, ref} from "vue";
import {httpGet, httpPost} from "@/utils/http";
import {ElMessage} from "element-plus";
import {dateFormat, disabledDate, removeArrayItem} from "@/utils/libs";
import {Plus} from "@element-plus/icons-vue";

// 变量定义
const items = ref([])
const item = ref({})
const showDialog = ref(false)
const rules = reactive({
  platform: [{required: true, message: '请选择平台', trigger: 'change',}],
  value: [{required: true, message: '请输入 API KEY 值', trigger: 'change',}]
})
const loading = ref(true)
const formRef = ref(null)
const title = ref("")
const platforms = ref(["Azure", "OpenAI", "ChatGLM"])

// 获取数据
httpGet('/api/admin/apikey/list').then((res) => {
  if (res.data) {
    // 初始化数据
    const arr = res.data;
    for (let i = 0; i < arr.length; i++) {
      arr[i].last_used_at = dateFormat(arr[i].last_used_at)
    }
    items.value = arr
  }
  loading.value = false
}).catch(() => {
  ElMessage.error("获取数据失败");
})

const add = function () {
  showDialog.value = true
  title.value = "新增 API KEY"
  item.value = {}
}

const edit = function (row) {
  showDialog.value = true
  title.value = "修改 API KEY"
  item.value = row
}

const save = function () {
  formRef.value.validate((valid) => {
    if (valid) {
      showDialog.value = false
      httpPost('/api/admin/apikey/save', item.value).then((res) => {
        ElMessage.success('操作成功！')
        if (!item.value['id']) {
          const newItem = res.data
          newItem.last_used_at = dateFormat(newItem.last_used_at)
          items.value.push(newItem)
        }
      }).catch((e) => {
        ElMessage.error('操作失败，' + e.message)
      })
    } else {
      return false
    }
  })
}

const remove = function (row) {
  httpGet('/api/admin/apikey/remove?id=' + row.id).then(() => {
    ElMessage.success("删除成功！")
    items.value = removeArrayItem(items.value, row, (v1, v2) => {
      return v1.id === v2.id
    })
  }).catch((e) => {
    ElMessage.error("删除失败：" + e.message)
  })
}
</script>

<style lang="stylus" scoped>
.list {

  .opt-box {
    padding-bottom: 10px;
    display flex;
    justify-content flex-end

    .el-icon {
      margin-right: 5px;
    }
  }

  .el-select {
    width: 100%
  }

}
</style>