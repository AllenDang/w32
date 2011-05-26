all: clean
	make -C pkg/w32			install
	make -C pkg/w32/kernel32	install
	make -C pkg/w32/user32		install
	make -C pkg/w32/gdi32		install
	make -C pkg/w32/comctl32	install
	make -C pkg/w32/comdlg32	install
	make -C pkg/w32/ole32		install
	make -C pkg/w32/shell32		install

clean:
	make -C pkg/w32			clean
	make -C pkg/w32/kernel32	clean
	make -C pkg/w32/user32		clean
	make -C pkg/w32/gdi32		clean
	make -C pkg/w32/comctl32	clean
	make -C pkg/w32/comdlg32	clean
	make -C pkg/w32/ole32		clean
	make -C pkg/w32/shell32		clean

format:
	gofmt -w .
