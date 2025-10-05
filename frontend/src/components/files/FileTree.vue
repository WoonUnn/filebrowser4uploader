<template>
	<div class="file-tree">
		<div v-if="error" class="error">{{ error }}</div>
		<ul v-else-if="rootNode" class="tree-root">
			<TreeNode :node="rootNode" :level="0" />
		</ul>
		<div v-else class="empty">暂无数据</div>
	</div>
</template>

<script setup>
import { ref, watch } from 'vue'
import TreeNode from './TreeNode.vue'
import { useFileStore } from '@/stores/file'

const fileStore = useFileStore()

const props = defineProps({
	// 传入的目录树 JSON 字符串
	treeDataString: { type: String, required: true }
})

const rootNode = ref(null)
const error = ref('')

function repairInput(str) {
	let s = String(str).trim()
	// 修复常见错误：右括号、data.children 起始应为数组
	s = s.replace(/\)/g, '}')
	const dataIdx = s.indexOf('"data"')
	if (dataIdx !== -1) {
		const after = s.slice(dataIdx)
		const m = after.match(/"children"\s*:\s*\{/)
		if (m) {
			const start = dataIdx + m.index
			s = s.slice(0, start) + s.slice(start).replace(/"children"\s*:\s*\{/, '"children":[{')
		}
	}
	return s
}

function normalizeNode(node) {
	const n = node && typeof node === 'object' ? node : {}
	const name = typeof n.name === 'string' ? n.name : 'root'
	const cv = n.children
	let children = []
	if (Array.isArray(cv)) {
		children = cv.map(normalizeNode)
	} else if (cv && typeof cv === 'object') {
		// 兼容 children 为对象（单子节点链）的情况
		children = [normalizeNode(cv)]
	}
	return { name, children, expanded: false }
}

function parseTreeString(str) {
	error.value = ''
	try {
		const fixed = repairInput(str)
		const obj = JSON.parse(fixed)
		const root = obj && obj.data ? obj.data : obj
		const normalized = normalizeNode(root)
		// 初始仅显示根目录
		normalized.expanded = true // Expand root by default
		return normalized
	} catch (e) {
		error.value = '目录数据解析失败，请提供合法的 JSON 字符串。'
		return null
	}
}

function build() {
	rootNode.value = parseTreeString(props.treeDataString)
}

watch(() => props.treeDataString, () => {
	build()
})
</script>

<style scoped>
.file-tree {
	font-family: system-ui, -apple-system, Segoe UI, Roboto, Arial, sans-serif;
	font-size: 14px;
	line-height: 1.4;
	padding: 8px;
}

.tree-root {
	padding: 0;
	margin: 0;
}

.error {
	color: #d93025;
	padding: 8px;
}

.empty {
	color: #777;
	padding: 8px;
}
</style>
