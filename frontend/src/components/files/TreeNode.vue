<template>
  <li>
    <div class="tree-row" :style="{ paddingLeft: (level * 20) + 'px' }" @click="onToggle">
      <span class="caret" :class="{ 'is-expanded': node.expanded }" v-if="hasChildren">
        <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
      </span>
      <span class="caret-placeholder" v-else></span>
      <span class="icon-folder">
        <i class="material-icons">folder</i>
      </span>
      <span class="folder-name">{{ node.name }}</span>
    </div>
    <ul v-show="node.expanded" v-if="hasChildren">
      <TreeNode v-for="(child, idx) in node.children"
                 :key="(child.name || 'node') + '-' + idx"
                 :node="child"
                 :level="level + 1" />
    </ul>
  </li>
</template>

<script setup>
import { computed } from 'vue';
// Recursive component needs to be imported.
// The name is implicitly set by the filename in <script setup>.
import TreeNode from './TreeNode.vue';

const props = defineProps({
  node: { type: Object, required: true },
  level: { type: Number, default: 0 }
});

const hasChildren = computed(() => Array.isArray(props.node.children) && props.node.children.length > 0);

const onToggle = (e) => {
  e.stopPropagation();
  if (hasChildren.value) {
    props.node.expanded = !props.node.expanded;
  }
};
</script>

<style lang="less" scoped>
li {
  list-style: none;
}
.tree-row {
  cursor: pointer;
  user-select: none;
  display: flex;
  align-items: center;
  white-space: nowrap;
  padding: 4px 8px;
  border-radius: 4px;
  transition: background-color 0.2s;
  .material-icons {
	color: #1d99f3;
  }
}
.tree-row:hover {
  background-color: #f0f0f0;
  .folder-name {
	color: #666;
  }
}
.caret {
  width: 16px;
  height: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
  margin-right: 4px;
  transition: transform 0.2s;
  transform: rotate(-90deg);
}
.caret.is-expanded {
  transform: rotate(0deg);
}
.caret-placeholder {
  width: 16px;
  margin-right: 4px;
}
.icon-folder {
  color: #888;
  margin-right: 8px;
  display: flex;
  align-items: center;
}
.folder-name {
  color: #fff;
  font-size: 15.5px;
  letter-spacing: 0.35px;
}
ul {
  margin: 0;
  padding: 0;
}
</style>
