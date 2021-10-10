// FINAL COMMANDS FOR NOW:
// Combine woman with border
ffmpeg -i ./in/woman.png -i ./in/border.png  -filter_complex "[1][0:v] overlay" ./in/womanWithBorder.png


// This output video + image
ffmpeg -i ./in/background.mp4 -i ./in/womanWithBorder.png \
-filter_complex "[0:v][1:v] overlay=80:0:enable='between(t,0,20)'" \
-pix_fmt yuv420p -b:v 2000k \
./in/outputBorderNoSound.mp4

// Add sound to video
ffmpeg -i ./in/outputBorderNoSound.mp4 -i ./in/squidgame.mp3 -map 0:v -map 1:a -c:v copy -shortest ./out/leakage-alert.mp4

// TEST COMMANDS:
//Outputs image with sound
// -acodec copy preserves the quality of the audio
// -r 1 i think is 1 repeat
// -shortest stops when one of the inputs comes to an end(mp3 in our case). Without -shortest, the processing wont stop
ffmpeg -loop 1 -i ./in/woman.png -i ./in/sound.mp3 -acodec copy -r 1 -shortest ./out/out.mp4

// This output video + image(closer to what we want)
ffmpeg -i ./in/background.mp4 -i ./in/woman.png \
-filter_complex "[0:v][1:v] overlay=100:0:enable='between(t,0,20)'" \
-pix_fmt yuv420p \
./out/output.mp4

//v2 - works but not all are layered together
ffmpeg \
-loop 1 -framerate 24 -t 10 -i ./in/woman.png \
-t 10 -i ./in/sound.mp3 \
-i ./in/background.mp4 \
-f lavfi -t 1 -i anullsrc \
-filter_complex \
"[0]setpts=PTS-STARTPTS,scale=1920x1080,fps=24,format=yuv420p[img1]; \
 [2]setpts=PTS-STARTPTS,scale=1920x1080,fps=24,format=yuv420p[vid1]; \
 [img1][1][vid1][3] concat=n=2:v=1:a=1" \
-vsync 2 \
-pix_fmt yuv420p \
./out/outputfile.mp4

ffmpeg \
-loop 1 -framerate 24 -t 10 -i ./in/woman.png \
-t 10 -i ./in/sound.mp3 \
-i ./in/background.mp4 \
-filter_complex \
"[0:v][2:v][1] overlay=80:0:enable='between(t,0,20)';\
-vsync 2 \
-pix_fmt yuv420p \
./out/outputfile1.mp4









