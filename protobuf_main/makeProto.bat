@echo off

echo start make proto
make proto

echo make proto over coolDown 10 second

TIMEOUT /T 10

pause