Problem: Too many re-renders when using useEffect.

Causes: Accidently setting the state of depedency array inside useEffect which cause infinite loop of useEffect being called.
