<script>
    import { LayoutDashboard, Settings, LogOut } from "@lucide/svelte";
    console.log("Left side bar")
    
</script>


<nav class="w-[100px] h-screen flex flex-col items-center gap-5 border-[#e2e8f0] border-r-2 p-5 fixed left-0 top-0">
    <div class="h-[40px] w-[40px] cursor-pointer">
        <img src="/logo.png" alt="GoodApps logo">
    </div>
    <div class="flex flex-col gap-4 mt-5 flex-grow">
        <div class="flex flex-col justify-center items-center gap-1 text-sm cursor-pointer">
            <LayoutDashboard size={20}/>
            <span class="font-bold text-xs">Dashboard</span>
        </div>
    </div>
    <div class="flex flex-col gap-4">
        <div class="flex flex-col justify-center items-center gap-1 text-sm cursor-pointer">
            <Settings  size={20}/>
            <span class="font-bold text-xs">Settings</span>
        </div>
        <div class="flex flex-col justify-center items-center gap-1 text-sm cursor-pointer">
            <LogOut size={20}/>
            <span class="font-bold text-xs">Logout</span>
        </div>
    </div>
</nav>
