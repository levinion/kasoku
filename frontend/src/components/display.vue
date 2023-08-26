<template>
    <v-main>
        <v-card class="card">
            <v-list>
                <v-virtual-scroll :items="search.list" height="72vh">
                    <template v-slot:default="item">
                        <!-- <v-tooltip  open-on-hover location="end"> -->
                            <!-- <template v-slot:activator="{ props: tip }"> -->
                                <v-list-item title="" @click="open(item.item)">
                                    <v-list-item-title :class="which_class(item.item)">
                                        {{ item.item.Name }}
                                    </v-list-item-title>
                                    <v-list-item-subtitle>
                                        {{ item.item.Path }}
                                    </v-list-item-subtitle>
                                </v-list-item>
                            <!-- </template> -->
                            <!-- <v-img v-if="search.isPicture(item.item)" :src="search.requireSrc(item.item.Path)" /> -->
                        <!-- </v-tooltip> -->
                    </template>
                </v-virtual-scroll>
            </v-list>
        </v-card>
    </v-main>
</template>

<script lang="ts" setup>
import { useSearchStore } from '../store/search';
import { FsInfo } from '../store/search';
import { Open } from '../../wailsjs/go/main/App';

const search = useSearchStore()
const open = (info: FsInfo) => {
    if (info.IsDir) {
        search.input = info.Path
        //显式触发搜索
        search.find()
    } else {
        Open(info.Path)
    }
}
function which_class(info: FsInfo): string {
    return info.IsDir ? "dir" : "file"
}
</script>

<style scoped>
.card {
    overflow: auto;
}

.dir {
    color: aqua
}

.file {
    color: chocolate
}
</style>