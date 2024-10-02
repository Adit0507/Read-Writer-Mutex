## Implementation of Read-Writer Mutex

<b><u>ReadLock()</u></b>: Synchronizes readers access using <i>readersLock</i>.

<b><u>WriteLock()</u></b>: Blocks writer's access for other goroutines using <i>globalLock</i>.

<b><u>ReadUnlock()</u></b>: Opens for other goroutine to gain access to the readers access by unlocking <i> readersLock. </i>

<b><u>WriteUnlock()</u></b>: Simply unlocks <i>globalLock</i> for the next goroutine to gain writing access.
