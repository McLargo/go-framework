package bootstrap

func RunCli() error {
	bootstrap, err := newBootstrap()
	if err != nil {
		return err
	}
	bootstrap.Logger.Info("Running cli...")
	return nil
}
