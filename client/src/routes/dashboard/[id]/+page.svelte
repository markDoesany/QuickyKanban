<script>
    import { CirclePlus } from "@lucide/svelte";
    import Headbar from "$lib/components/Headbar.svelte";
    import KanbanCard from "$lib/components/KanbanCard.svelte";
</script>
 
<div class="h-screen w-screen p-5">
    <div class="bg-[#F8FAFC] flex flex-col gap-8">
        <h1 class="font-bold text-2xl">Project X</h1>
        <h2 class="text-[#475569] font-semibold text-lg">Task List</h2>
        <ul class="flex gap-5 text-[#475569] font-semibold">
            <li class="hover:cursor-pointer border-b-2 border-[#0aadc2] p-2">All</li>
            <li class="hover:cursor-pointer border-b-2 border-transparent hover:border-[#1E293B] p-2">To Do</li>
            <li class="hover:cursor-pointer border-b-2 border-transparent hover:border-[#1E293B] p-2">In Progress</li>
            <li class="hover:cursor-pointer border-b-2 border-transparent hover:border-[#1E293B] p-2">Completed</li>
        </ul>
    </div>
    <main class="h-full w-full mt-10 grid grid-cols-3 gap-4 px-4">
        <section class="flex flex-col items-center gap-5 bg-[#F8FAFC] p-4 rounded-2xl shadow-md h-fit">
            <!-- <Headbar color="#4F46E5" title="To Do" count={20}/> -->
             <div class="w-full flex justify-between items-center">
                 <div class="flex items-center gap-2 w-full font-bold text-lg">
                    <div class="w-2 h-2 bg-[#4F46E5] rounded-full"></div>
                    <h2>To Do</h2>
                    <span class="text-[#94A3B8]">(8)</span>
                 </div>
                 <CirclePlus color="#475569" size={25} class="cursor-pointer"/>
             </div>
            <div class="w-full space-y-3">
                <KanbanCard />
            </div>
        </section>
    
        <section class="flex flex-col items-center gap-5 bg-[#F8FAFC] p-4 rounded-lg shadow-md h-fit">
            <Headbar color="#F59E0B" title="In Progress" count={2}/>
            <div class="w-full space-y-3">
                <KanbanCard />
                <KanbanCard />
                <KanbanCard />
                <KanbanCard />
                <KanbanCard />
                <KanbanCard />
                <KanbanCard />
                <KanbanCard />
                <KanbanCard />
            </div>
        </section>
    
        <section class="flex flex-col items-center gap-5 bg-[#F8FAFC] p-4 rounded-lg shadow-md h-fit">
            <Headbar color="#22C55E" title="Completed" count={40}/>
            <div class="w-full space-y-3">
                <KanbanCard />
                <KanbanCard />
            </div>
        </section>
    </main>
    
</div>