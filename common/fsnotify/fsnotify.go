package ggfsnotify

import (
	"errors"
	log "github.com/alecthomas/log4go"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
)

type WatchFile struct {
	watch *fsnotify.Watcher
	f     func(file string)
	dir   string
	done  chan struct{}
}

func init() {

}

func NewWatchFile() (*WatchFile, error) {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err)
		return nil, errors.New("create watcher failed")
	}

	w := &WatchFile{watch: watch, done: make(chan struct{}, 1)}

	return w, nil
}

func (wf *WatchFile) watchDir(dir string) error {
	if wf == nil {
		return errors.New("object instance is empty")
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err = filepath.Abs(path)
			if err != nil {
				return err
			}

			err = wf.watch.Add(path)
			if err != nil {
				return err
			}
			//log.Info("monitor: %v", path)
		}
		return nil
	})

	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case evs := <-wf.watch.Events:
				{
					if evs.Op&fsnotify.Create == fsnotify.Create {
						log.Info("Create: %v", evs.Name)
						var fi os.FileInfo
						fi, err = os.Stat(evs.Name)
						if err == nil && fi.IsDir() {
							err = wf.watch.Add(evs.Name)
							if err != nil {

							}
							log.Info("Create monitor: %s", evs.Name)
						}
					}

					if evs.Op&fsnotify.Write == fsnotify.Write {
						log.Info("Write: %s", evs.Name)
						wf.f(evs.Name)
					}

					if evs.Op&fsnotify.Remove == fsnotify.Remove {
						log.Info("Remove: %s", evs.Name)
						var fi os.FileInfo
						fi, err = os.Stat(evs.Name)
						if err == nil && fi.IsDir() {
							err = wf.watch.Remove(evs.Name)
							if err != nil {

							}
							log.Info("Remove monitor: %s", evs.Name)
						}
					}

					if evs.Op&fsnotify.Rename == fsnotify.Rename {
						err = wf.watch.Remove(evs.Name)
						if err != nil {
							log.Error(err)
						}
						log.Info("Rename: %s", evs.Name)
					}

					if evs.Op&fsnotify.Chmod == fsnotify.Chmod {
						log.Info("Chmod: %s", evs.Name)
					}
				}
			case err = <-wf.watch.Errors:
				{
					if err != nil {
						log.Error("error: %v", err)
					}
					return
				}
			case <-wf.done:
				{
					err = wf.watch.Close()
					if err != nil {
						log.Error("error: %v", err)
					}
					log.Info("monitor file change exit")
				}
			}
		}
	}()

	return nil
}

func (wf *WatchFile) Start(dir string, f func(file string)) error {
	if wf == nil {
		return errors.New("object instance is empty")
	}

	wf.f = f
	wf.dir = dir
	er := wf.watchDir(wf.dir)
	if er != nil {
		return er
	}

	return nil
}

func (wf *WatchFile) Stop() error {
	if wf == nil {
		return errors.New("object instance is empty")
	}

	if wf.done != nil {
		wf.done <- struct{}{}
	}

	return nil
}
