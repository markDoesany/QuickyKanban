<script>
    import Headbar from "$lib/components/Headbar.svelte";
    import KanbanCard from "$lib/components/KanbanCard.svelte";
</script>
 
<div class="h-screen w-screen p-5 overflow-hidden ">
    <p>Kanban Dashboard</p>
    <div class="h-full w-full">
        <Headbar/>
        <div class="w-[250px]">
            <KanbanCard/>
        </div>
    </div>
</div>