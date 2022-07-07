import {h} from "vue";
import {NIcon} from "naive-ui";

export function renderIcon(icon: any) {
    return () => h(NIcon, null, {default: () => h(icon)})
}
