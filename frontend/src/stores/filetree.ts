import { defineStore } from "pinia";

export const useFiletreeStore = defineStore("filetree", {
	state: () => ({
		treeDataString: "", // 用于存储文件目录树的字符串表示
	}),
	actions: {
		updateTreeData(newData: string) {
			this.treeDataString = newData;
		},
		clearTreeData() {
			this.treeDataString = "";
		}
	},
});