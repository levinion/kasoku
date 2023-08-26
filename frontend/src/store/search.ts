import { defineStore } from 'pinia'
import { AbsoluteSearch, BackToLastDir, GetBase64ImageSrc } from "../../wailsjs/go/main/App"
import { FuseSearch } from "../../wailsjs/go/main/App"
import { GetCurrentDir } from "../../wailsjs/go/main/App"
export const useSearchStore = defineStore('search', {
    state: () => ({
        input: "",
        fuseInput: "",
        jumpHiddenItem: true,
        loading: false,
        list: [] as FsInfo[],
        picture: "",
    }),
    actions: {
        find() {
            if (this.input === "") {
                return
            }
            AbsoluteSearch(this.input, this.jumpHiddenItem).catch(err => { console.log(err) }).then(r => {
                if (r) {
                    this.list = r
                }
            })
        },
        fuseFind() {
            if (this.fuseInput === "") {
                return
            }
            this.loading = true
            FuseSearch(this.fuseInput, this.jumpHiddenItem).catch(err => { console.log(err) }).then(r => {
                if (r) {
                    this.list = r
                }
                this.loading = false
            })
        },
        clearFuseInput() {
            this.fuseInput = ""
        },
        initInput() {
            GetCurrentDir().then(res => {
                this.input = res
            }).then(() => {
                this.find()
            })
        },
        toggleJumpHiddenItem() {
            this.jumpHiddenItem = !this.jumpHiddenItem
            this.find()
        },
        backToLastDir() {
            BackToLastDir().then(res => {
                this.input = res
            }).then(() => {
                this.find()
            })
        },
        isPicture(info: FsInfo): boolean {
            if (!info.IsDir) {
                let path = info.Path.toLowerCase()
                if (path.endsWith(".jpeg") || path.endsWith(".jpg") || path.endsWith(".png") || path.endsWith(".gif") || path.endsWith(".bmp")) {
                    return true
                }
            }
            return false
        },
        getPreviewPictureBase64(path: string) {
            GetBase64ImageSrc(path).then(res => {
                this.picture = res
            })
        },
        requireSrc(path: string) {
            this.getPreviewPictureBase64(path)
            return this.picture
        }
    }
})

export interface FsInfo {
    Name: string,
    Path: string,
    IsDir: boolean,
}