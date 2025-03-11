<script>
    import { LayoutDashboard, CalendarFold, Settings, LogOut } from "@lucide/svelte";
    console.log("Left side bar")
    
</script>


<nav class="w-[120px] min-h-screen flex flex-col items-center gap-5 border-r-gray-400 border p-5">
    <div class="h-[50px] w-[50px] cursor-pointer">
        <img src="/logo.png" alt="GoodApps logo">
    </div>
    <div class="flex flex-col gap-4 mt-5 flex-grow">
        <div class="flex flex-col justify-center items-center gap-1 text-sm cursor-pointer">
            <LayoutDashboard/>
            <span>Dashboard</span>
        </div>
        <div class="flex flex-col justify-center items-center gap-1 text-sm cursor-pointer">
            <CalendarFold />
            <span>Calendar</span>
        </div>
    </div>
    <div class="flex flex-col gap-4">
        <div class="flex flex-col justify-center items-center gap-1 text-sm cursor-pointer">
            <Settings />
            <span>Settings</span>
        </div>
        <div class="flex flex-col justify-center items-center gap-1 text-sm cursor-pointer">
            <LogOut />
            <span>Logout</span>
        </div>
    </div>
</nav>
