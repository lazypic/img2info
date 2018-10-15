# img2str
tesseract 테스트 리포지터리입니다. 아직까지 좋은 결과가 나오진 않습니다.
단순한 이미지에는 사용할 수 있으며, 좋은 결과를 위해서는 
tesseract4를 기다려봐야할듯합니다.
gosseract보다는 tesseract가 현재는 결과가 더 좋습니다.

#### How to use
- 이미지에 들어간 텍스트 정보를 가지고 옵니다.
```
$ brew install tesseract
$ tesseract blender_opengl_test.png out
```
- input
![blender_opengl_test](https://user-images.githubusercontent.com/1149996/46948833-50080f00-d0ba-11e8-96b7-142243740586.png)

- output
```
File /Users/kimhanwoong/Desktop/mammaiopenngest7v004.blend Strip <none>
Date 2018/10/15 19:48:59
RenderTime 00:25.21

Peak Memory 293.39M
openGL image



)

Marker <none> Timecode 00:00:00:17 Frame 017 Camera cam01 Lens 24.00 Scene Scene
```
#### Gosseract Install
- 실행하기 위해서는 tesseract가 설치되어 있어야 합니다.
```
$ brew install tesseract
$ go get -t github.com/otiai10/gosseract
```

#### Reference
- https://github.com/tesseract-ocr/tesseract
- https://github.com/tesseract-ocr/tesseract/wiki
- https://github.com/otiai10/gosseract
