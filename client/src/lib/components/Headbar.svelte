<script lang="ts">
    import {Plus} from "@lucide/svelte";
    let {title="Title Here", color="#4F46E5", count=0}:{title:string, color:string, count:number} = $props()
</script>

<button class="rounded-full text-white text-sm font-semibold w-full h-[35px] flex items-center justify-between p-2 cursor-pointer transition-all duration-200 hover:bg-[#1E293B] active:scale-95 shadow-md"
        style="background-color: {color};">
    <div class="flex items-center gap-2">
        <span class="bg-white text-[${color}] rounded-full w-3 flex items-center justify-center text-xs px-4 py-1"
              style="color:{color};">
                {count}</span>
        <span class="text-sm">{title}</span>
    </div>
    <span class="font-bold w-7 h-7 flex items-center justify-center rounded-full">
        <Plus />
    </span>
</button>
